package uuid

import (
	"net/http"
	"testing"
)

func TestParseGenerator(t *testing.T) {
	t.Run("valid generator", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/uuid?v=v5", nil)

		version, _, _ := parseGenerator(request)
		expectedVersion := "v5"

		if version != expectedVersion {
			t.Errorf("Expected version %s, got %s", expectedVersion, version)
		}
	})

	t.Run("missing generator", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/uuid", nil)

		version, _, _ := parseGenerator(request)
		expectedVersion := "v4"

		if version != expectedVersion {
			t.Errorf("Expected version %s, got %s", expectedVersion, version)
		}
	})

	t.Run("invalid generator", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/uuid?v=v3", nil)

		_, _, err := parseGenerator(request)

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestParseN(t *testing.T) {
	t.Run("valid n", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/uuid?n=5", nil)

		parsed, _ := parseN(request)
		expected := 5

		if parsed != expected {
			t.Errorf("Expected %d, got %d", expected, parsed)
		}
	})

	t.Run("invalid n", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/uuid?n=hello", nil)

		_, err := parseN(request)

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

	t.Run("out-of-range n", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/uuid?n=99999", nil)

		_, err := parseN(request)

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

	t.Run("missing n", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/uuid", nil)

		parsed, _ := parseN(request)
		expected := 1

		if parsed != expected {
			t.Errorf("Expected %d, got %d", expected, parsed)
		}
	})
}
