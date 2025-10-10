package authenticate

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal(err)
	}
}

func TestAuthenticate(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// All sorts of things will be going on here
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
      "access_token": "eyJhbGciOiJSUzI1NiIsImtpZCI6IjI3QjA1ODk4Nzc1OEUwMkMI1NiIsInR2...",
      "expires_in": 3600,
      "token_type": "Bearer",
      "scope": "public_api_shipmentserver"
    }`))
	}))
	// Close the server
	defer mockServer.Close()

	ctx := context.Background()

	// Establish the config struct
	cfg := Config{
		ClientId:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
	}
	var err error

	if _, err = Authenticate(ctx, mockServer.URL, cfg); err != nil {
		t.Fatalf("Authenticate failed %v", err)
	}
}
