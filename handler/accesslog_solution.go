package handler

import (
	"encoding/json"
	"github.com/go-sre/copilot/accesslog"
	"github.com/go-sre/core/exchange"
	"github.com/go-sre/core/runtime"
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
					status = runtime.NewStatus(http.StatusInternalServerError, location, err)
				} else {
					exchange.WriteResponse(w, buf, status)
				}
			}
		}
	case http.MethodDelete:
		accesslog.Delete()
	case http.MethodPut:
		var entry accesslog.Entry
		entry, status = exchange.Deserialize[runtime.LogError, accesslog.Entry](nil, r.Body)
		if status.OK() {
			accesslog.Put(entry)
		}
	default:
	}
	exchange.WriteResponse(w, nil, status)
}
