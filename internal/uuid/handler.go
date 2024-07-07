package uuid

import (
	"encoding/json"
	"log"
	"net/http"
)

// Handler is a route handler for generating UUIDs
func Handler(w http.ResponseWriter, r *http.Request) {
	// Attempt to parse the generator function and UUID spec version from the request
	version, generator, err := parseGenerator(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Attempt to parse `n` from the request (number of UUIDs to generate)
	n, err := parseN(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Generate `n` UUIDs
	uuids, err := generateUuids(generator, n)
	if err != nil {
		log.Printf("Failed to generate UUID (version %s): %s\n", version, err)
		http.Error(w, "Failed to generate UUIDs!", http.StatusInternalServerError)
		return
	}

	// Marshal generated UUIDs into JSON response
	resp, err := json.Marshal(uuids)
	if err != nil {
		log.Printf("Failed to marshal UUID response JSON data: %s\n", err)
		http.Error(w, "Failed to generate UUIDs!", http.StatusInternalServerError)
		return
	}

	// Write JSON response to http.ResponseWriter
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(resp); err != nil {
		log.Printf("Failed to write to HTTP response writer: %s\n", err)
		http.Error(w, "Failed to generate UUIDs!", http.StatusInternalServerError)
	}
}
