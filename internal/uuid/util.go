package uuid

import (
	"errors"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

var uuidGenerators = map[string]func() (uuid.UUID, error){
	"v1": uuid.NewUUID,
	"v4": uuid.NewRandom,
	"v6": uuid.NewV6,
	"v7": uuid.NewV7,
}

func parseGenerator(r *http.Request) (string, func() (uuid.UUID, error), error) {
	version := r.PathValue("version")

	generator, ok := uuidGenerators[version]
	if !ok {
		return version, nil, errors.New("invalid UUID version parameter")
	}

	return version, generator, nil
}

func parseN(r *http.Request) (int, error) {
	n := 1

	if r.URL.Query().Has("n") {
		newN, err := strconv.Atoi(r.URL.Query().Get("n"))
		if err != nil {
			return 0, err
		}

		if newN < 1 || newN > 1000 {
			return 0, errors.New("query parameter 'n' must be between 1 and 1000")
		}

		n = newN
	}

	return n, nil
}
