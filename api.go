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
    // Your code here

    // Send an HTTP response
    io.WriteString(w, "OK")
}