package uuid

import "github.com/google/uuid"

// generateUuids produces `n` UUIDS using the `generator` function, using goroutines for concurrency, and collecting
// all generated UUIDs into a single slice of strings.
func generateUuids(generator func() (uuid.UUID, error), n int) ([]string, error) {
	uuids := make([]string, n)
	errChan := make(chan error, n)
	uuidChan := make(chan string, n)

	for i := 0; i < n; i++ {
		// Spawn a goroutine to generate the UUID
		go func() {
			generated, err := generator()
			if err != nil {
				// Send the error (err) through the errors channel (errChan)
				errChan <- err
				return
			}
			// Send the generated UUID as a string through the uuids channel (uuidChan)
			uuidChan <- generated.String()
		}()
	}

	// If any of the generators resulted in an error, return the error
	for i := 0; i < n; i++ {
		select {
		case err := <-errChan:
			return nil, err
		case generated := <-uuidChan:
			uuids[i] = generated
		}
	}

	// Return the generated UUIDs
	return uuids, nil
}
