package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	log.Fatal(http.ListenAndServe(":"+port, http.HandlerFunc(indexHandler)))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello pod!\n"))
}
