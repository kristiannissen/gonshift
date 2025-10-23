package addons

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"gonshift/internal/restclient"
	"gonshift/pkg/models"
)

func GetDropPoint(ctx context.Context, cfg *models.Config) (*models.DataObject, error) {
	return &models.DataObject{}, nil
}

func GetDropPoints(ctx context.Context, cfg *models.Config, payload *models.DataObject) (*models.DataObject, error) {
	p, _ := json.Marshal(payload)

	c := restclient.NewRestClient(
		restclient.WithEndpoint(cfg.Endpoint),
		restclient.WithAccessToken(cfg.AccessToken),
	)

	if b, err := c.Post(ctx, p); err != nil {
		log.Println(err)
		return &models.DataObject{}, err
	} else {
		//
		d := &models.DataObject{}
		if err := json.NewDecoder(bytes.NewReader(b)).Decode(&d.Data); err != nil {
			log.Println(err)
			return d, err
		}
		return d, nil
	}
}
