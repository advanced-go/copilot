package handler

import (
	"encoding/json"
	"github.com/advanced-go/copilot/accesslog"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/runtime"
	"net/http"
)

func AccessLogHandler_Solution(w http.ResponseWriter, r *http.Request) {
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
					http2.WriteResponse(w, buf, status, nil)
					return
				}
			} else {
				status = runtime.NewStatusError(http.StatusNotFound, location, nil)
			}
		}
	case http.MethodDelete:
		accesslog.Delete()
	case http.MethodPut:
		var entry accesslog.Entry
		entry, status = http2.Deserialize[accesslog.Entry](r.Body)
		if status.OK() {
			accesslog.Put(entry)
		}
	default:
	}
	http2.WriteResponse(w, nil, status, nil)
}
