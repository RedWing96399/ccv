package server

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

// server implementation is still incomplete
func run() {

	var addr string
	flag.StringVar(&addr, "addr", "http://localhost:8080", "address flag")

	s := http.Server{
		Addr:        addr,
		Handler:     nil,
		IdleTimeout: time.Minute * 1,
	}

	err := s.ListenAndServe()
	if err != nil {
		fmt.Printf("server error : %v", err.Error())
	}

}
