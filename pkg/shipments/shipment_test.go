package shipments

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path"
	"testing"

	"gonshift/testhelper"
)

func TestGetDocumentLists(t *testing.T) {
	mockServer := testhelper.NewTestServer(t, http.StatusOK, testhelper.Fixture(t, "getdocumentlist_response.json"))
	// Close the server
	defer mockServer.Close()

	ctx := context.Background()

	cfg := Config{
		AccessToken: os.Getenv("ACCESS_TOKEN"),
		ActorId:     os.Getenv("ACTOR_ID"),
		Endpoint:    mockServer.URL,
	}

	s := map[string]any{}
	json.Unmarshal(testhelper.Fixture(t, "getdocumentlists_request.json"), &s)

	if _, err := GetDocumentLists(ctx, cfg, DataObject{
		Data: s,
	}); err != nil {
		t.Error(err)
	}
}

func TestSaveShipment(t *testing.T) {
	mockServer := testhelper.NewTestServer(t, http.StatusOK, testhelper.Fixture(t, "shipment_response.json"))
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
	mockServer := testhelper.NewTestServer(t, http.StatusOK, testhelper.Fixture(t, "documents.json"))
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
	mockServer := testhelper.NewTestServer(t, http.StatusOK, testhelper.Fixture(t, "shipments.json"))
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
