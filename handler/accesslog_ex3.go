package handler

import (
	"encoding/json"
	"github.com/advanced-go/copilot/accesslog"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/io2"
	"github.com/advanced-go/core/runtime"
	"net/http"
)

func AccessLogHandler_3(w http.ResponseWriter, r *http.Request) {
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
					http2.WriteResponse(w, buf, status,nil)
				}
			}
		}
	case http.MethodDelete:
		accesslog.Delete()
	case http.MethodPut:
		var entry accesslog.Entry

		// Start unmarshalling
		buf, err := io2.ReadAll(r.Body)
		//github:copilot
		if err != nil {
			status = runtime.NewStatusError(http.StatusInternalServerError, location, err)
		} else {
			err = json.Unmarshal(buf, &entry)
			if err != nil {
				status = runtime.NewStatusError(http.StatusInternalServerError, location, err)
			} else {
				accesslog.Put(entry)
			}
		}
	default:
	}
	http2.WriteResponse(w, nil, status,nil)
}

// Refactoring by just typeing func and waiting
func AccessLogHandler_4(w http.ResponseWriter, r *http.Request) {
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
					http2.WriteResponse(w, buf, status,nil)
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

		// Start unmarshalling
		buf, err := io2.ReadAll(r.Body)
		//github:copilot
		if err != nil {
			status = runtime.NewStatusError(http.StatusInternalServerError, location, err)
		} else {
			err = json.Unmarshal(buf, &entry)
			if err != nil {
				status = runtime.NewStatusError(http.StatusInternalServerError, location, err)
			} else {
				accesslog.Put(entry)
			}
		}
	default:
	}
	http2.WriteResponse(w, nil, status,nil)
}

// Start a function for refactoring case http.MethodPut from above, entering the function name.
func unmarshal
