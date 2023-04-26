package handler

import (
	"github.com/go-sre/ai/percept"
	"github.com/go-sre/core/exchange"
	"github.com/go-sre/core/runtime"
	"net/http"
)

func AgentHandler_Ex1_Copilot(w http.ResponseWriter, r *http.Request) {
	var status = runtime.NewStatusOK()

	switch r.Method {
	case http.MethodPost:
		// The exercise is to finish the code for the POST method
		var o percept.Observation

		buf, err := exchange.ReadAll(r.Body)
		if err != nil {
			status = runtime.NewStatusBadRequest(err)
		}
	case http.MethodDelete:
	case http.MethodPut:

	default:
	}
	exchange.WriteResponse(w, nil, status)
}
