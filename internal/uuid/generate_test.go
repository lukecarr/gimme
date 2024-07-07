package uuid

import (
	"errors"
	"github.com/google/uuid"
	"testing"
)

func TestGenerateUuids(t *testing.T) {
	t.Run("generates n UUIDs", func(t *testing.T) {
		expected := 25
		uuids, _ := generateUuids(uuid.NewRandom, expected)

		if len(uuids) != expected {
			t.Errorf("Expected %d UUIDs, got %d\n", expected, len(uuids))
		}
	})

	t.Run("handles errors", func(t *testing.T) {
		_, err := generateUuids(func() (uuid.UUID, error) {
			return uuid.New(), errors.New("dummy error")
		}, 10)

		if err == nil {
			t.Errorf("Expected error, got nil\n")
		}
	})
}
