package shipments

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"gonshift/internal/restclient"
	"gonshift/pkg/models"
)

// GetShipments is used to return a shipment which has not been deleted
// Example usage:
//
//	s := GetShipments(context.Content &models.Config) // OUTPUT
func GetShipments(ctx context.Context, cfg *models.Config) (*models.DataObject, error) {
	c := restclient.NewRestClient(
		restclient.WithEndpoint(cfg.Endpoint),
		restclient.WithAccessToken(cfg.AccessToken),
	)
	// This method has no body
	if b, err := c.Get(ctx, []byte(``)); err != nil {
		log.Println(err)
		return &models.DataObject{}, err
	} else {
		d := &models.DataObject{}
		json.Unmarshal(b, &d.Data)
		return d, nil
	}
}

// GetDocumentLists returns a list of all documents on a shipment
// Example usage:
func GetDocumentLists(ctx context.Context, cfg *models.Config, dataobject *models.DataObject) (*models.DataObject, error) {
	c := restclient.NewRestClient(
		restclient.WithEndpoint(cfg.Endpoint),
		restclient.WithAccessToken(cfg.AccessToken),
	)

	s, err := json.Marshal(dataobject)
	if err != nil {
		return &models.DataObject{}, err
	}

	if b, err := c.Get(ctx, s); err != nil {
		log.Println(err)
		return &models.DataObject{}, err
	} else {
		d := &models.DataObject{}
		if err := json.NewDecoder(bytes.NewReader(b)).Decode(&d.Data); err != nil {
			log.Println(err)
			return &models.DataObject{}, err
		}
		return d, nil
	}
}

// GetDocuments returns a list of all documents encoded as base64 on a shipment.
// Example usage:
func GetDocuments(ctx context.Context, cfg *models.Config) (*models.DataObject, error) {
	c := restclient.NewRestClient(
		restclient.WithEndpoint(cfg.Endpoint),
		restclient.WithAccessToken(cfg.AccessToken),
	)

	if b, err := c.Get(ctx, []byte(``)); err != nil {
		log.Println(err)
		return &models.DataObject{}, err
	} else {
		// Decode documents
		d := &models.DataObject{}
		if err := json.NewDecoder(bytes.NewReader(b)).Decode(&d.Data); err != nil {
			log.Println(err)
			return &models.DataObject{}, nil
		}
		return d, nil
	}
}

// SaveShipment is used to save a shipment to Inbox
// Example usage:
func SaveShipment(ctx context.Context, cfg *models.Config, dataobject *models.DataObject) (*models.DataObject, error) {
	c := restclient.NewRestClient(
		restclient.WithEndpoint(cfg.Endpoint),
		restclient.WithAccessToken(cfg.AccessToken),
	)
	// Marshal the shipment
	s, err := json.Marshal(dataobject)
	if err != nil {
		return &models.DataObject{}, err
	}

	// This method has no body
	if b, err := c.Post(ctx, s); err != nil {
		log.Println(err)
		return &models.DataObject{}, err
	} else {
		d := &models.DataObject{}
		json.Unmarshal(b, &d.Data)
		return d, nil
	}
}
