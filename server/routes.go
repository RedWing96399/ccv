package server

import (
	"net/http"
)

func getRoutes() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", validate)
	mux.HandleFunc("/a", valC)

	return mux
}
