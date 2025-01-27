package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func getRequiredEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Required environment variable %s is not set ", key)
	}
	return value
}
func main() {
	// Optional env var with default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// Required env var
	secretMessage := getRequiredEnv("SECRET_MESSAGE")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World on 27th January!\nSecret message: %s", secretMessage)
	})
	fmt.Printf("Listening on port after revuild %s\n", port)
	http.ListenAndServe(":"+port, nil)
}

