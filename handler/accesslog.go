package handler

import (
	"encoding/json"
	"github.com/advanced-go/copilot/accesslog"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
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
					status = runtime.NewStatusError(http.StatusInternalServerError, location, err)
				} else {
					http2.WriteResponse[runtime.LogError](w, buf, status, nil)
					return
				}
			} else {
				status = runtime.NewStatusError(http.StatusNotFound, location, nil)
			}
		}
	case http.MethodDelete:
		accesslog.Delete()
	case http.MethodPut:
		// The exercise is to finish the code for the PUT method
	default:
	}
	http2.WriteResponse(w, nil, status, nil)
}
