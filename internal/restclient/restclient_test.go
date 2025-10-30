package restclient

//

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"gonshift/testhelper"
)

func TestRestClientPostSuccess(t *testing.T) {
	mockServer := testhelper.NewTestServer(t, http.StatusOK, []byte(`{"message": "OK"}`))

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

	mockServer := testhelper.NewTestServer(t, http.StatusUnauthorized, testhelper.Fixture(t, "errors.json"))

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
	)

	if _, err := c.Get(ctx, []byte(``)); err == nil {
		t.Errorf("error %s", err)
	}

	if ctx.Err() != context.DeadlineExceeded {
		t.Errorf("error %s", ctx.Err())
	}

}
