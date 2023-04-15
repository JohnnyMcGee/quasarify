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

// Function myHTTPFunction is an HTTP handler
func myHTTPFunction(w http.ResponseWriter, r *http.Request) {
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
