package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/token", tokenHandler)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)

	println(err)
}