package uuid

import (
	"errors"
	"net/http"
	"strconv"
)

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
