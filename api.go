package api

import (
	"io"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	// Register an HTTP function with the Functions Framework
	functions.HTTP("api", myHTTPFunction)
}

func myHTTPFunction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	message := r.URL.Query().Get("message")
	qMessage := ""
	if message != "" {
		qMessage = QuasarifyMessage(message)
	} else {
		qMessage = "Transmission received... in space..."
	}
	// Send an HTTP response
	io.WriteString(w, qMessage)
}
