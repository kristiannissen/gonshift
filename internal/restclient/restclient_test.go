package restclient

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(path.Join("..", "..", ".env")); err != nil {
		log.Fatal(err)
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
		WithDebug(true),
		WithEndpoint(mockServer.URL),
		WithAccessToken(os.Getenv("ACCESS_TOKEN")),
	)

	if b, err := c.Post(ctx, []byte(`{"hello": "Kitty"}`)); err != nil {
		t.Errorf("error %s", err)
	} else {
		log.Println(string(b))
	}
}
