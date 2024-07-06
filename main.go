package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func hello(w http.ResponseWriter, _ *http.Request) {
	if _, err := fmt.Fprintf(w, "Hello, World!\n"); err != nil {
		log.Fatalf("Failed to write to http.ResponseWriter: %s\n", err)
	}
}

func main() {
	http.HandleFunc("/", hello)

	addr := os.Getenv("ADDR")
	if strings.TrimSpace(addr) == "" {
		addr = ":5000"
	}

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Failed to start HTTP server on '%s': %s\n", addr, err)
	}
}
