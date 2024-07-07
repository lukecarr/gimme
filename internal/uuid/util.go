package uuid

import (
	"errors"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

// uuidGenerators is a map of UUID spec versions to functions that produce UUIDs (and errors)
var uuidGenerators = map[string]func() (uuid.UUID, error){
	"v1": uuid.NewUUID,
	"v4": uuid.NewRandom,
	"v6": uuid.NewV6,
	"v7": uuid.NewV7,
}

// parseGenerator attempts to extract a `version` from the request's query parameters and then find the
// corresponding UUID generator function from the `uuidGenerators` map
func parseGenerator(r *http.Request) (string, func() (uuid.UUID, error), error) {
	if r.URL.Query().Has("v") {
		version := r.URL.Query().Get("v")

		generator, ok := uuidGenerators[version]
		if !ok {
			return version, nil, errors.New("invalid UUID version parameter")
		}

		return version, generator, nil
	}

	return "v4", uuidGenerators["v4"], nil
}

// parseN attempts to extract an integer between 1 and 1000 from the request's query parameters
func parseN(r *http.Request) (int, error) {
	if r.URL.Query().Has("n") {
		n, err := strconv.Atoi(r.URL.Query().Get("n"))
		if err != nil {
			return 0, err
		}

		if n < 1 || n > 1000 {
			return 0, errors.New("query parameter 'n' must be between 1 and 1000")
		}

		return n, nil
	}

	return 1, nil
}
