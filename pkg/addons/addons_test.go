package addons

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"gonshift/pkg/models"
	"gonshift/testhelper"
)

func TestGetDropPoints(t *testing.T) {
	mockServer := testhelper.NewTestServer(t, http.StatusOK, testhelper.Fixture(t, "addons/droppoints_response.json"))
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
	json.Unmarshal(testhelper.Fixture(t, "addons/droppoints_request.json"), &s)

	// Payload
	payload := &models.DataObject{
		Data:    s,
		Options: map[string]any{},
	}

	if _, err := GetDropPoints(ctx, cfg, payload); err != nil {
		t.Fatalf("DropPoints failed %v", err)
	}
}
