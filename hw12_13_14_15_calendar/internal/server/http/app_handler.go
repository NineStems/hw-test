package internalhttp

import (
	"net/http"

	"github.com/hw-test/hw12_13_14_15_calendar/common"
)

func (s *ServerHTTP) health() func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	return func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		w.Write([]byte("health is ok, action-id=" + r.Context().Value(common.CtxActionID).(string)))
	}
}
