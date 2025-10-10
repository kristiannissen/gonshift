package shipments

import (
	"context"
	"encoding/json"
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
		log.Println(err)
	}
}

// TODO: Move to a reusable helper function
// arguments would be *testing, filename
//
//	[]byte
func fixture(t *testing.T, name string) []byte {
	// Helper marks the calling function as a test helper function
	t.Helper()
	// Load fixture file and return []byte()
	if d, err := os.ReadFile(path.Join("..", "..", "fixtures", name)); err != nil {
		log.Println(err)
		return []byte(``)
	} else {
		return d
	}
}

func TestSaveShipment(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// All sorts of things will be going on here
		w.WriteHeader(http.StatusOK)
		// Return the fixture data
		w.Write(fixture(t, "shipment_response.json"))
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
	// Load and marshal shipment to json
	s := map[string]any{}
	if b, err := os.ReadFile(path.Join("..", "..", "fixtures", "shipment_request.json")); err != nil {
		log.Println(err)
	} else {
		json.Unmarshal(b, &s)
	}

	if _, err := SaveShipment(ctx, cfg, DataObject{
		Data:    s,
		Options: map[string]any{"Labels": "PDF"},
	}); err != nil {
		t.Error(err)
	}
}

func TestGetDocuments(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// All sorts of things will be going on here
		w.WriteHeader(http.StatusOK)
		// Return the fixture data
		w.Write(fixture(t, "documents.json"))
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

	if _, err := GetShipments(ctx, cfg); err != nil {
		t.Error(err)
	}
}

func TestGetShipments(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// All sorts of things will be going on here
		w.WriteHeader(http.StatusOK)
		// Return the fixture data
		w.Write(fixture(t, "shipments.json"))
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

	if _, err := GetShipments(ctx, cfg); err != nil {
		t.Error(err)
	}
}
