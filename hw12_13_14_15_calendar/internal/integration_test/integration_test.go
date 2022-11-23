//go:build integration
// +build integration

package integrations_test

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	v1 "github.com/hw-test/hw12_13_14_15_calendar/api/v1"
)

func TestUseCase(t *testing.T) {
	url := "http://localhost:8888"

	events := []string{`{
    "title": "title",
    "ownerId": 1,
    "date": "2022-11-15T16:00:00+00:00",
    "date_end": "2022-11-15T17:00:00+00:00",
    "date_notification": "2022-11-15T15:59:00+00:00",
    "description": "description"
	}`, `{
    "title": "title",
    "ownerId": 1,
    "date": "2022-11-14T16:00:00+00:00",
    "date_end": "2022-11-14T17:00:00+00:00",
    "date_notification": "2022-11-14T15:59:00+00:00",
    "description": "description"
	}`, `{
    "title": "title",
    "ownerId": 1,
    "date": "2022-11-02T16:00:00+00:00",
    "date_end": "2022-11-02T17:00:00+00:00",
    "date_notification": "2022-11-02T15:59:00+00:00",
    "description": "description"
	}`,
	}

	var ids []string
	defer func(linkIds *[]string) {
		for _, lids := range *linkIds {
			req, _ := http.NewRequest(http.MethodDelete, url+"/event/"+lids+"/delete", nil)
			_, _ = http.DefaultClient.Do(req)
		}
	}(&ids)

	t.Run("Успешное добавление события", func(t *testing.T) {
		for _, event := range events {
			eventReader := strings.NewReader(event)
			req, err := http.NewRequest(http.MethodPost, url+"/event/create", eventReader)
			if err != nil {
				t.Errorf("problem with creating request: %v", err)
				return
			}

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Errorf("client: error making http request: %v", err)
				return
			}

			if res.StatusCode != 200 {
				t.Errorf("client: wrong status code: %v", res.StatusCode)
				return
			}

			resBody, err := io.ReadAll(res.Body)
			if err != nil {
				t.Errorf("client: could not read response body: %s\n", err)
				return
			}
			resp := v1.CreateResponse{}
			err = json.Unmarshal(resBody, &resp)
			if err != nil {
				t.Errorf("Unmarshal body: %s\n", err)
				return
			}
			ids = append(ids, resp.Id)
		}

	})

	t.Run("Ошибка при повторном добавление события", func(t *testing.T) {
		for _, event := range events {
			eventReader := strings.NewReader(event)
			req, err := http.NewRequest(http.MethodPost, url+"/event/create", eventReader)
			if err != nil {
				t.Errorf("problem with creating request: %v", err)
				return
			}

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Errorf("client: error making http request: %v", err)
				return
			}

			if res.StatusCode != 500 {
				t.Errorf("client: wrong status code: %v", res.StatusCode)
				return
			}
		}
	})

	t.Run("Получение событий за день", func(t *testing.T) {
		dayRequest := `{
			"date": "2022-11-15T22:00:00Z",
			"condition": 1
		}`
		reader := strings.NewReader(dayRequest)
		req, err := http.NewRequest(http.MethodPost, url+"/event/read", reader)
		if err != nil {
			t.Errorf("problem with creating request: %v", err)
			return
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Errorf("client: error making http request: %v", err)
			return
		}

		if res.StatusCode != 200 {
			t.Errorf("client: wrong status code: %v", res.StatusCode)
			return
		}

		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			t.Errorf("client: could not read response body: %s\n", err)
			return
		}

		var resp v1.ReadResult
		err = json.Unmarshal(resBody, &resp)
		if err != nil {
			t.Errorf("Unmarshal body: %s\n", err)
			return
		}

		if len(resp.Events) != 1 {
			t.Errorf("wrong len of resp want 1 got %d\n", len(resp.Events))
			return
		}
	})

	t.Run("Получение событий за неделю", func(t *testing.T) {
		weekRequest := `{
			"date": "2022-11-15T22:00:00Z",
			"condition": 2
		}`
		reader := strings.NewReader(weekRequest)
		req, err := http.NewRequest(http.MethodPost, url+"/event/read", reader)
		if err != nil {
			t.Errorf("problem with creating request: %v", err)
			return
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Errorf("client: error making http request: %v", err)
			return
		}

		if res.StatusCode != 200 {
			t.Errorf("client: wrong status code: %v", res.StatusCode)
			return
		}

		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			t.Errorf("client: could not read response body: %s\n", err)
			return
		}

		var resp v1.ReadResult
		err = json.Unmarshal(resBody, &resp)
		if err != nil {
			t.Errorf("Unmarshal body: %s\n", err)
			return
		}

		if len(resp.Events) != 2 {
			t.Errorf("wrong len of resp want 2 got %d\n", len(resp.Events))
			return
		}
	})

	t.Run("Получение событий за месяц", func(t *testing.T) {
		monthRequest := `{
			"date": "2022-11-15T22:00:00Z",
			"condition": 3
		}`
		reader := strings.NewReader(monthRequest)
		req, err := http.NewRequest(http.MethodPost, url+"/event/read", reader)
		if err != nil {
			t.Errorf("problem with creating request: %v", err)
			return
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Errorf("client: error making http request: %v", err)
			return
		}

		if res.StatusCode != 200 {
			t.Errorf("client: wrong status code: %v", res.StatusCode)
			return
		}

		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			t.Errorf("client: could not read response body: %s\n", err)
			return
		}

		var resp v1.ReadResult
		err = json.Unmarshal(resBody, &resp)
		if err != nil {
			t.Errorf("Unmarshal body: %s\n", err)
			return
		}

		if len(resp.Events) != 3 {
			t.Errorf("wrong len of resp want 3 got %d\n", len(resp.Events))
			return
		}
	})

	//os.Exit(0)
}
