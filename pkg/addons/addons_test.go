package addons

import (
	"context"
	"net/http"
	"os"
	"testing"

	"gonshift/testhelper"
)

func TestGetDropPoint(t *testing.T) {
	mockServer := testhelper.NewTestServer(t, http.StatusOK, testhelper.Fixture(t, "droppoints.json"))
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
