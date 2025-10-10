package restclient

//

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(path.Join("..", "..", ".env")); err != nil {
		log.Fatal(err)
	}
}

func fixture(t *testing.T, name string) []byte {
	// Helper marks the calling function as a test helper function
	t.Helper()
	// Load fixture file and return []byte()
	if d, err := os.ReadFile(path.Join("..", "..", "fixtures", name)); err != nil {
		log.Fatal(err)
		return []byte(``)
	} else {
		return d
	}
}

func TestRestClientPostSuccess(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		// Write response body
		w.Write([]byte(`{"message": "OK"}`))
	}))
	// Close the server
	defer mockServer.Close()

	ctx := context.Background()

	c := NewRestClient(
		WithDebug(false),
		WithEndpoint(mockServer.URL),
		WithAccessToken(os.Getenv("ACCESS_TOKEN")),
	)

	if _, err := c.Post(ctx, []byte(`{"hello": "Kitty"}`)); err != nil {
		t.Errorf("error %s", err)
	}
}

func TestRestClientError(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		// Write response body
		w.Write(fixture(t, "errors.json"))
	}))
	// Close the server
	defer mockServer.Close()

	ctx := context.Background()

	c := NewRestClient(
		WithEndpoint(mockServer.URL),
		WithAccessToken(os.Getenv("ACCESS_TOKEN")),
		WithDebug(false),
	)

	if _, err := c.Get(ctx, []byte(`{"hello": "Kitty"}`)); err.Error() != "Unauthorized" {
		t.Errorf("error %s", err)
	}
}

func TestRestClientTimeout(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create slow response
		time.Sleep(200 * time.Millisecond)
		// Provide status
		w.WriteHeader(http.StatusOK)
		// Write response body
		w.Write([]byte(``))
	}))
	// Close the server
	defer mockServer.Close()

	ctx, cancel := context.WithTimeout(context.Background(), (50 * time.Millisecond))
	defer cancel()

	c := NewRestClient(
		WithEndpoint(mockServer.URL),
		WithAccessToken(os.Getenv("ACCESS_TOKEN")),
		WithDebug(true),
	)

	if _, err := c.Get(ctx, []byte(``)); err == nil {
		t.Errorf("error %s", err)
	}

	if ctx.Err() != context.DeadlineExceeded {
		t.Errorf("error %s", ctx.Err())
	}

}
