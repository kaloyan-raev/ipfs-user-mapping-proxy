package mock

import "net/http"

// ErrorHandler is an HTTP handler that always responds with an Internal Server Error.
func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "error", http.StatusInternalServerError)
}
