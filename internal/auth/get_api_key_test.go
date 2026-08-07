package auth

import(
	"testing"
	"net/http"
)

func TestGetAPIKey_NoHeader(t *testing.T) {
	headers := http.Header{}

	_, err := GetAPIKey(headers)

	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf("expected %v, got %v", ErrNoAuthHeaderIncluded, err)
	}
}

func TestGetAPIKey_MalformedHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer abc123")

	_, err := GetAPIKey(headers)

	if err == nil {
		t.Fatal("expected an error")
	}
}

func TestGetAPIKey_Success(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey abc123")

	key, err := GetAPIKey(headers)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if key != "abc123" {
		t.Errorf("expected %q, got %q", "abc123", key)
	}
}
