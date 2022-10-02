package internalhttp

import (
	"context"
	"log"
	"net/http"

	"github.com/calendar/hw12_13_14_15_calendar/common"
	"github.com/calendar/hw12_13_14_15_calendar/internal/pkg/util"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				log.Fatal(rvr)
			}
		}()

		actionID, err := util.GenerateUUID()
		if err != nil {
			panic(err)
		}
		ctx := context.WithValue(context.Background(), common.CtxActionID, actionID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
		// todo
	})
}
