package index

import (
	"net/http"
)

// Handle `/` endpoint
func Handle(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("Hello World")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
