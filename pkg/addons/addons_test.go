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

func setupTestDataObject(t *testing.T) []byte {
	// Helper marks the calling function as a test helper function
	t.Helper()
	// Load fixture file and return []byte()
	if d, err := os.ReadFile(path.Join("..", "..", "fixtures", "droppoints.json")); err != nil {
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
		w.Write(setupTestDataObject(t))
	}))
	// Close the server
	defer mockServer.Close()

	ctx := context.Background()

	// Create the config struct
	cfg := Config{
		Access_Token: os.Getenv("ACCESS_TOKEN"),
		ActorId:      os.Getenv("ACTOR_ID"),
	}

	// Payload
	payload := DataObject{
		Data: Data{
			Kind:          1,
			ProdConceptID: 553,
			Services:      []int{47001, 47023},
			Addresses: []Address{
				{
					Kind:        1,
					Name1:       "Test Test",
					Street1:     "Test Street 1",
					PostCode:    "7224",
					City:        "Verdal",
					CountryCode: "NO",
				},
			},
		},
		Options: map[string]any{},
	}

	if _, err := GetDropPoints(ctx, mockServer.URL+"/%s", cfg, payload); err != nil {
		t.Fatalf("DropPoints failed %v", err)
	}
}
