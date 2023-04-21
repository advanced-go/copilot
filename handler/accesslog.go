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
				}
			}
		}
	case http.MethodDelete:
		accesslog.Delete()
	case http.MethodPut:
		var entry accesslog.Entry

		// First attempt
		buf, err := exchange.ReadAll(r.Body)
		// Copilot created this snippet - which is correct, but is too verbose. Although, it does use the
		// runtime.Status type and accesslog.Put() function correctly, so it has learned that.
		if err != nil {
			status = runtime.NewStatus(http.StatusInternalServerError, location, err)
		} else {
			err = json.Unmarshal(buf, &entry)
			if err != nil {
				status = runtime.NewStatus(http.StatusInternalServerError, location, err)
			} else {
				accesslog.Put(entry)
			}
		}

		// Second attempt - copilot implemented the unmarshal() function that was started
		entry, status = unmarshal(r)

		// Correct solution - copilot did not implement or even know about this. Not in its training data
		entry, status = exchange.Deserialize[runtime.LogError, accesslog.Entry](nil, r.Body)
		if status.OK() {
			accesslog.Put(entry)
		}
	default:
	}
	exchange.WriteResponse(w, nil, status)
}

//github:copilot - exchange.HandleFunc() is not a function, even intellisense knows that
// I was trying to create the unmarshal() function, and copilot created the function init()
/*
func init() {
	exchange.HandleFunc("/access-log", AccessLogHandler)
} // Path: pkg\handler\accesslog_test.go	// Path: pkg\handler\accesslog.go	// Path: pkg\accesslog\access_test.go	// Path: pkg\host\startup.go

*/

//github:copilot - copilot started the func init() in a comment
//github:copilot <[func init() {]> - <[copilot created the func init() in a comment]> - <[copilot started the func init() in a comment]> - <[copilot started the func init() in a comment]> 	 	  // Path: pkg\handler\accesslog_test.go	// Path: pkg\handler\accesslog.go	// Path: pkg\accesslog\access_test.go	// Path: pkg\host\startup.go  	 	   // Path: pkg\handler\accesslog_test.go	// Path: pkg\handler\accesslog.go	// Path: pkg\accesslog\access_test.go	// Path: pkg\host\startup.go
// Copilot did Ok here but does not seem to understand pointers, as the return Status type <[should be a pointer.]>
// I find it amusing that Copilot advice was used to complete the above sentence, so Copilot can recommend
// text that defines the issue, pointer vs non-pointer, but cannot recommend code that understands the issue.
// Probably this is due to the training data <[not having enough examples of pointers.]>
//
func unmarshal(r *http.Request) (accesslog.Entry, runtime.Status) {
	var entry accesslog.Entry
	buf, err := exchange.ReadAll(r.Body)
	if err != nil {
		return entry, runtime.NewStatus(http.StatusInternalServerError, location, err)
	}
	err = json.Unmarshal(buf, &entry)
	if err != nil {
		return entry, runtime.NewStatus(http.StatusInternalServerError, location, err)
	}
	return entry, runtime.NewStatusOK()
}
