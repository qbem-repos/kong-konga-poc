package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(os.Getenv("PORT"), nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, fmt.Sprint("{\"service-name\":\"", os.Getenv("NAME"), "\"}"))
}
