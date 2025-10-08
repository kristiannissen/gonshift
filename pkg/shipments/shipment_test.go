package shipments

import (
	"context"
	"fmt"
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

// TODO: Move to a reusable helper function
// arguments would be *testing, filename
//
//	[]byte
func setupTestDataObject(t *testing.T) []byte {
	// Helper marks the calling function as a test helper function
	t.Helper()
	// Load fixture file and return []byte()
	if d, err := os.ReadFile(path.Join("..", "..", "fixtures", "shipment.json")); err != nil {
		log.Fatal(err)
		return []byte(``)
	} else {
		return d
	}
}

func TestGetShipmentSuccess(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// All sorts of things will be going on here
		w.WriteHeader(http.StatusOK)
		// Return the fixture data
		w.Write(setupTestDataObject(t))
	}))
	// Close the server
	defer mockServer.Close()

	ctx := context.Background()

	// Create the config struct
	cfg := Config{
		AccessToken: os.Getenv("ACCESS_TOKEN"),
		ActorId:     os.Getenv("ACTOR_ID"),
		Endpoint:    fmt.Sprintf(mockServer.URL, os.Getenv("ACTOR_ID")),
	}

	if _, err := GetShipment(ctx, cfg); err != nil {
		t.Error(err)
	}
}
