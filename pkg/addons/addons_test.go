package addons

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

func TestGetDropPoint(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// All sorts of things will be going on here
		w.WriteHeader(http.StatusOK)
		// Return the fixture data
		w.Write(fixture(t, "droppoints.json"))
	}))
	// Close the server
	defer mockServer.Close()

	ctx := context.Background()

	// Create the config struct
	cfg := Config{
		AccessToken: os.Getenv("ACCESS_TOKEN"),
		ActorId:     os.Getenv("ACTOR_ID"),
		Endpoint:    mockServer.URL,
	}

	// Payload
	payload := DataObject{
		Data:    map[string]any{},
		Options: map[string]any{},
	}

	if _, err := GetDropPoints(ctx, cfg, payload); err != nil {
		t.Fatalf("DropPoints failed %v", err)
	}
}
