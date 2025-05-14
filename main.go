package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, GitHub App!")
	})

	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received a webhook event!")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Webhook received")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8094"
	}
	log.Printf("Listening on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
