package handler
import (
	"encoding/json"
	"github.com/go-sre/ai/percept"
	"github.com/go-sre/core/exchange"
	"github.com/go-sre/core/runtime"
	"net/http"
)


func AgentHandler_Ex4(w http.ResponseWriter, r *http.Request) {
	var status = runtime.NewStatusOK()

	switch r.Method {
	case http.MethodPost:
		// The exercise is to finish the code for the POST method
		var o percept.Observation

		o, status = unmarshal(r)

	case http.MethodDelete:
	case http.MethodPut:

	default:
	}
	exchange.WriteResponse(w, nil, status)
}

func unmarshal(r *http.Request



