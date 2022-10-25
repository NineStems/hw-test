package server

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	v1 "github.com/hw-test/hw12_13_14_15_calendar/api/v1"
	logmock "github.com/hw-test/hw12_13_14_15_calendar/common/mocks"
	"github.com/hw-test/hw12_13_14_15_calendar/internal/config"
	"github.com/hw-test/hw12_13_14_15_calendar/internal/server/mocks"
)

func TestServer(t *testing.T) {
	ctx := context.Background()
	cfg := config.Config{
		Server: config.Server{
			Grpc: config.Grpc{
				Host: "127.0.0.1",
				Port: "12201",
			},
			Http: config.Rest{
				Host: "127.0.0.1",
				Port: "8080",
			},
		},
		Database: config.Database{
			Source: "inmemory",
		},
	}

	srv := NewServer(logmock.MockLogger{}, &cfg, &mocks.MockApp{})
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go srv.Start(ctx)
	time.Sleep(time.Millisecond * 1)

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(":12201", opts...)
	require.NoError(t, err)

	defer conn.Close()
	client := v1.NewCalendarClient(conn)

	t.Run("create grpc", func(t *testing.T) {
		mes, err := client.CreateEvent(context.Background(), &v1.CreateRequest{
			OwnerId:          1,
			Title:            "test",
			Date:             "2006-01-02T15:04:05Z",
			DateEnd:          "2006-01-02T15:04:05Z",
			DateNotification: "2006-01-02T15:04:05Z",
			Description:      "test",
		})
		require.NoError(t, err)
		require.NotEmpty(t, mes.Id)
	})

	t.Run("update grpc", func(t *testing.T) {
		_, err = client.UpdateEvent(ctx, &v1.UpdateRequest{
			Id:               "test",
			OwnerId:          1,
			Title:            "test",
			Date:             "2006-01-02T15:04:05Z",
			DateEnd:          "2006-01-02T15:04:05Z",
			DateNotification: "2006-01-02T15:04:05Z",
			Description:      "test",
		})
		require.NoError(t, err)
	})

	t.Run("delete grpc", func(t *testing.T) {
		_, err = client.DeleteEvent(ctx, &v1.DeleteRequest{
			Id: "test",
		})
		require.NoError(t, err)
	})

	t.Run("read grpc", func(t *testing.T) {
		res, err := client.ReadEvents(ctx, &v1.ReadRequest{
			Date:      time.Now().Format(time.RFC3339),
			Condition: 0,
		})
		require.NoError(t, err)
		require.Len(t, res.Events, 2)
	})

	restClient := http.Client{}
	t.Run("create rest", func(t *testing.T) {
		event := v1.CreateRequest{
			OwnerId:          1,
			Title:            "test",
			Date:             "2006-01-02T15:04:05Z",
			DateEnd:          "2006-01-02T15:04:05Z",
			DateNotification: "2006-01-02T15:04:05Z",
			Description:      "test",
		}
		body, err := json.Marshal(&event)
		reader := bytes.NewReader(body)
		require.NoError(t, err)
		req, err := http.NewRequest("POST", "http://127.0.0.1:8080/event/create", reader)
		require.NoError(t, err)

		resp, err := restClient.Do(req)
		require.NoError(t, err)
		respbody, err := io.ReadAll(resp.Body)
		resp.Body.Close()

		var respData v1.CreateResponse
		err = json.Unmarshal(respbody, &respData)
		require.NoError(t, err)
		require.NotEmpty(t, respData.Id)
	})

	t.Run("update rest", func(t *testing.T) {
		event := v1.UpdateRequest{
			OwnerId:          1,
			Title:            "test",
			Date:             "2006-01-02T15:04:05Z",
			DateEnd:          "2006-01-02T15:04:05Z",
			DateNotification: "2006-01-02T15:04:05Z",
			Description:      "test",
		}
		body, err := json.Marshal(&event)
		reader := bytes.NewReader(body)
		require.NoError(t, err)
		req, err := http.NewRequest("PUT", "http://127.0.0.1:8080/event/test/update", reader)
		require.NoError(t, err)

		_, err = restClient.Do(req)
		require.NoError(t, err)
	})

	t.Run("delete rest", func(t *testing.T) {
		req, err := http.NewRequest("PUT", "http://127.0.0.1:8080/event/test/delete", nil)
		require.NoError(t, err)

		_, err = restClient.Do(req)
		require.NoError(t, err)
	})

	t.Run("read rest", func(t *testing.T) {
		event := v1.ReadRequest{
			Date:      "2006-01-02T15:04:05Z",
			Condition: 0,
		}
		body, err := json.Marshal(&event)
		reader := bytes.NewReader(body)
		require.NoError(t, err)
		req, err := http.NewRequest("POST", "http://127.0.0.1:8080/event/read", reader)
		require.NoError(t, err)

		resp, err := restClient.Do(req)
		require.NoError(t, err)
		respbody, err := io.ReadAll(resp.Body)
		resp.Body.Close()

		var respData v1.ReadResult
		err = json.Unmarshal(respbody, &respData)
		require.NoError(t, err)
		require.Len(t, respData.Events, 2)
	})
}
