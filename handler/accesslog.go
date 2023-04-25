package handler

import (
	"encoding/json"
	"github.com/go-sre/copilot/accesslog"
	"github.com/go-sre/core/exchange"
	"github.com/go-sre/core/runtime"
	"net/http"
)

var location = "/handler/access-log"

func AccessLogHandler(w http.ResponseWriter, r *http.Request) {
	var status = runtime.NewStatusOK()

	switch r.Method {
	case http.MethodGet:
		if r.URL.Query() != nil && r.URL.Query().Get("delete") == "true" {
			accesslog.Delete()
		} else {
			entries := accesslog.Get()
			if len(entries) > 0 {
				buf, err := json.Marshal(entries)
				if err != nil {
					status = runtime.NewStatus(http.StatusInternalServerError, location, err)
				} else {
					exchange.WriteResponse(w, buf, status)
					return
				}
			} else {
				status = runtime.NewStatus(http.StatusNotFound, location, nil)
			}
		}
	case http.MethodDelete:
		accesslog.Delete()
	case http.MethodPut:
		// The exercise is to finish the code for the PUT method
	default:
	}
	exchange.WriteResponse(w, nil, status)
}
