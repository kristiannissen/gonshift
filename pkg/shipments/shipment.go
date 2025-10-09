package shipments

import (
	"context"
	"encoding/json"
	"log"

	"gonshift/internal/restclient"
)

func init() {

}

type Config struct {
	ActorId     string
	AccessToken string
	Endpoint    string
}

type DataObject struct {
	Data    any `json:"data"`
	Options any `json:"options,omitempty"`
}

func GetShipment(ctx context.Context, cfg Config) (any, error) {
	c := restclient.NewRestClient(
		restclient.WithEndpoint(cfg.Endpoint),
		restclient.WithAccessToken(cfg.AccessToken),
	)
	// This method has no body
	if b, err := c.Get(ctx, []byte(``)); err != nil {
		log.Fatal(err)
		return map[string]any{}, err
	} else {
		s := map[string]any{}
		json.Unmarshal(b, &s)
		return s, nil
	}
}

func SaveShipment(ctx context.Context, cfg Config, dataobject DataObject) (any, error) {
	c := restclient.NewRestClient(
		restclient.WithEndpoint(cfg.Endpoint),
		restclient.WithAccessToken(cfg.AccessToken),
	)
	// Marshal the shipment
	s, err := json.Marshal(dataobject)
	if err != nil {
		return map[string]any{}, err
	}

	// This method has no body
	if b, err := c.Post(ctx, s); err != nil {
		log.Fatal(err)
		return map[string]any{}, err
	} else {
		s := map[string]any{}
		json.Unmarshal(b, &s)
		return s, nil
	}
}
