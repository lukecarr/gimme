package main

import (
	"github.com/lukecarr/gimme/internal/uuid"
	"log"
	"net/http"
	"os"
	"strings"
)

// main is the entrypoint for the program
func main() {
	// Register UUID route
	http.HandleFunc("/uuid/{version}", uuid.Handler)

	// Attempt to parse address from environment variable
	addr := os.Getenv("ADDR")
	// Set default address value if not found in environment
	if strings.TrimSpace(addr) == "" {
		addr = ":5000"
	}

	// Start HTTP server and listen for incoming requests
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Failed to start HTTP server on '%s': %s\n", addr, err)
	}
}
