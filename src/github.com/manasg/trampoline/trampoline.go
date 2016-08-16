package main

import (
	"github.com/gorilla/handlers"
	"io"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "yo")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
}
