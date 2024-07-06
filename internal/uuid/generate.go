package uuid

import "github.com/google/uuid"

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
