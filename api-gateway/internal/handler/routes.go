package handler

import (
	"fmt"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/example", exampleHandler)
}

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from /example route")
}
