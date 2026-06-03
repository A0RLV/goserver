package main

import (
	"net/http"
	"fmt"
)

func main() {
	mux := http.NewServeMux()
	server := http.Server{
		Addr: ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	fmt.Errorf("%v", err)
}
