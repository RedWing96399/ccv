package server

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

// server implementation is still incomplete
func Run() {

	var addr string
	flag.StringVar(&addr, "addr", "localhost:8080", "address flag")

	mux := getRoutes()
	s := http.Server{
		Addr:        addr,
		Handler:     mux,
		IdleTimeout: time.Minute * 1,
	}

	err := s.ListenAndServe()
	if err != nil {
		fmt.Printf("server error : %v", err.Error())
	}

}
