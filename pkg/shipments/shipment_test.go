package shipments

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"gonshift/pkg/models"
	"gonshift/testhelper"
)

func TestGetDocumentLists(t *testing.T) {
	mockServer := testhelper.NewTestServer(t, http.StatusOK, testhelper.Fixture(t, "shipments/getdocumentlist_response.json"))
	// Close the server
	defer mockServer.Close()

	ctx := context.Background()

	cfg := &models.Config{
		AccessToken: os.Getenv("ACCESS_TOKEN"),
		ActorId:     os.Getenv("ACTOR_ID"),
		Endpoint:    mockServer.URL,
	}

	s := map[string]any{}
	json.Unmarshal(testhelper.Fixture(t, "shipments/getdocumentlists_request.json"), &s)

	if _, err := GetDocumentLists(ctx, cfg, &models.DataObject{
		Data: s,
	}); err != nil {
		t.Error(err)
	}
}

func TestSaveShipment(t *testing.T) {
	mockServer := testhelper.NewTestServer(t, http.StatusOK, testhelper.Fixture(t, "shipments/shipment_response.json"))
	// Close the server
	defer mockServer.Close()

	ctx := context.Background()

	// Create the config struct
	cfg := &models.Config{
		AccessToken: os.Getenv("ACCESS_TOKEN"),
		ActorId:     os.Getenv("ACTOR_ID"),
		Endpoint:    mockServer.URL,
	}
	// Load and marshal shipment to json
	s := map[string]any{}
	json.Unmarshal(testhelper.Fixture(t, "shipments/shipment_request.json"), &s)

	if _, err := SaveShipment(ctx, cfg, &models.DataObject{
		Data:    s,
		Options: map[string]any{"Labels": "PDF"},
	}); err != nil {
		t.Error(err)
	}
}

func TestGetDocuments(t *testing.T) {
	mockServer := testhelper.NewTestServer(t, http.StatusOK, testhelper.Fixture(t, "shipments/documents.json"))
	// Close the server
	defer mockServer.Close()

	ctx := context.Background()

	// Create the config struct
	cfg := &models.Config{
		AccessToken: os.Getenv("ACCESS_TOKEN"),
		ActorId:     os.Getenv("ACTOR_ID"),
		Endpoint:    mockServer.URL,
	}

	if _, err := GetShipments(ctx, cfg); err != nil {
		t.Error(err)
	}
}

func TestShipments(t *testing.T) {
	mockServer := testhelper.NewTestServer(t, http.StatusOK, testhelper.Fixture(t, "shipments/shipments_response.json"))
	// Close the server
	defer mockServer.Close()

	ctx := context.Background()

	// Create the config struct
	cfg := &models.Config{
		AccessToken: os.Getenv("ACCESS_TOKEN"),
		ActorId:     os.Getenv("ACTOR_ID"),
		Endpoint:    mockServer.URL,
	}

	if _, err := Shipments(ctx, cfg); err != nil {
		t.Error(err)
	}
}
