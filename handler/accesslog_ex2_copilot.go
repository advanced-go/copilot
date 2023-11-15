package handler

import (
	"encoding/json"
	"github.com/advanced-go/copilot/accesslog"
	"github.com/advanced-go/core/http2"
	"github.com/advanced-go/core/io2"
	"github.com/advanced-go/core/runtime"
	"net/http"
)

func AccessLogHandler_2_Copilot(w http.ResponseWriter, r *http.Request) {
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
	http2.WriteResponse(w, nil, status)
}

//github:copilot
func init() {
	http2.HandleFunc("/access-log", AccessLogHandler)
} // Path: pkg\handler\accesslog_test.go	// Path: pkg\handler\accesslog.go	// Path: pkg\accesslog\access_test.go	// Path: pkg\host\startup.go

//github:copilot func init() {

//github:copilot
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
					http2.WriteResponse(w, buf, status)
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
	http2.WriteResponse(w, nil, status, nil)
}
