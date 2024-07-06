package uuid

import (
	"encoding/json"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	version, generator, err := parseGenerator(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	n, err := parseN(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	uuids, err := generateUuids(generator, n)
	if err != nil {
		log.Printf("Failed to generate UUID (version %s): %s\n", version, err)
		http.Error(w, "Failed to generate UUIDs!", http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(uuids)
	if err != nil {
		log.Printf("Failed to marshal UUID response JSON data: %s\n", err)
		http.Error(w, "Failed to generate UUIDs!", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(resp); err != nil {
		log.Printf("Failed to write to HTTP response writer: %s\n", err)
		http.Error(w, "Failed to generate UUIDs!", http.StatusInternalServerError)
	}
}
