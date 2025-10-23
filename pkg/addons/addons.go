package addons

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"gonshift/internal/restclient"
)

type Config struct {
	ActorId     string
	AccessToken string
	Endpoint    string
}

type DataObject struct {
	Data    map[string]interface{} `json:"data"`
	Options map[string]interface{} `json:"options,omitempty"`
}

func GetDropPoint(ctx context.Context, cfg Config) (DataObject, error) {
	return DataObject{}, nil
}

func GetDropPoints(ctx context.Context, cfg Config, payload DataObject) (DataObject, error) {
	p, _ := json.Marshal(payload)

	c := restclient.NewRestClient(
		restclient.WithEndpoint(cfg.Endpoint),
		restclient.WithAccessToken(cfg.AccessToken),
	)

	if b, err := c.Post(ctx, p); err != nil {
		log.Println(err)
		return DataObject{}, err
	} else {
		//
		var d DataObject
		if err := json.NewDecoder(bytes.NewReader(b)).Decode(&d.Data); err != nil {
			log.Println(err)
			return d, err
		}
		return d, nil
	}
}
