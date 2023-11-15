package handler

import (
	"encoding/json"
	"github.com/advanced-go/copilot/accesslog"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/io2"
	"github.com/advanced-go/core/runtime"
	"net/http"
)

func AccessLogHandler_4_Copilot(w http.ResponseWriter, r *http.Request) {
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

		// Refactoring unmarshalling
		entry, status := unmarshal(r)

	default:
	}
	http2.WriteResponse(w, nil, status, nil)
}

//github:copilot
func unmarshal(r *http.Request) (accesslog.Entry, runtime.Status) {
	var entry accesslog.Entry
	buf, err := io2.ReadAll(r.Body)
	if err != nil {
		return entry, runtime.NewStatusError(http.StatusInternalServerError, location, err)
	}
	err = json.Unmarshal(buf, &entry)
	if err != nil {
		return entry, runtime.NewStatusError(http.StatusInternalServerError, location, err)
	}
	return entry, runtime.NewStatusOK()
}
