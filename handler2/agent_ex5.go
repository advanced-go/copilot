package handler

import (
	"encoding/json"
	"github.com/go-sre/ai/actuator"
	"github.com/go-sre/ai/percept"
	"github.com/go-sre/core/exchange"
	"github.com/go-sre/core/runtime"
	"net/http"
	"net/url"
)

func AgentHandler_Ex5(w http.ResponseWriter, r *http.Request) {
	var status = runtime.NewStatusOK()

	switch r.Method {
	case http.MethodPost:
		// The exercise is to finish the code for the POST method
		var o percept.Observation

		o, status = unmarshal(r)
		if status.OK() {
			signal("origin", url.Values{}, actuator.HttpActuator{})
		}
	case http.MethodDelete:
	case http.MethodPut:

	default:
	}
	exchange.WriteResponse(w, nil, status)
}

func unmarshal(r *http.Request) (percept.Observation, runtime.Status) {
	var o percept.Observation
	var status = runtime.NewStatusOK()
	buf, err := exchange.ReadAll(r.Body)
	if err != nil {
		status = runtime.NewStatus(http.StatusInternalServerError, location, err)
	} else {
		err = json.Unmarshal(buf, &o)
		if err != nil {
			status = runtime.NewStatus(http.StatusInternalServerError, location, err)
		}
	}
	return o, status
}
