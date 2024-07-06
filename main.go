package main

import (
	"github.com/lukecarr/gimme/internal/uuid"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/uuid/{version}", uuid.Handler)

	addr := os.Getenv("ADDR")
	if strings.TrimSpace(addr) == "" {
		addr = ":5000"
	}

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Failed to start HTTP server on '%s': %s\n", addr, err)
	}
}
