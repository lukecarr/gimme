package uuid

import "github.com/google/uuid"

var uuidGenerators = map[string]func() (uuid.UUID, error){
	"v1": uuid.NewUUID,
	"v4": uuid.NewRandom,
	"v6": uuid.NewV6,
	"v7": uuid.NewV7,
}

func generateUuids(generator func() (uuid.UUID, error), n int) ([]string, error) {
	uuids := make([]string, n)
	errChan := make(chan error, n)
	uuidChan := make(chan string, n)

	for i := 0; i < n; i++ {
		go func() {
			generated, err := generator()
			if err != nil {
				errChan <- err
				return
			}
			uuidChan <- generated.String()
		}()
	}

	for i := 0; i < n; i++ {
		select {
		case err := <-errChan:
			return nil, err
		case generated := <-uuidChan:
			uuids[i] = generated
		}
	}

	return uuids, nil
}
