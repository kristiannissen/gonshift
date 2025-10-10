package shipments

import (
	"bytes"
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
	Data    map[string]interface{} `json:"data"`
	Options map[string]interface{} `json:"options,omitempty"`
}

// https://helpcenter.nshift.com/hc/en-us/articles/4401176907676-API-Documentation#get_shipments
// Used to return a shipment which has not been deleted.
func GetShipments(ctx context.Context, cfg Config) (DataObject, error) {
	c := restclient.NewRestClient(
		restclient.WithEndpoint(cfg.Endpoint),
		restclient.WithAccessToken(cfg.AccessToken),
	)
	// This method has no body
	if b, err := c.Get(ctx, []byte(``)); err != nil {
		log.Println(err)
		return DataObject{}, err
	} else {
		var do DataObject
		json.Unmarshal(b, &do.Data)
		return do, nil
	}
}

func GetDocuments(ctx context.Context, cfg Config) (DataObject, error) {
	c := restclient.NewRestClient(
		restclient.WithEndpoint(cfg.Endpoint),
		restclient.WithAccessToken(cfg.AccessToken),
	)

	if b, err := c.Get(ctx, []byte(``)); err != nil {
		log.Println(err)
		return DataObject{}, err
	} else {
		// Decode documents
		var d DataObject
		if err := json.NewDecoder(bytes.NewReader(b)).Decode(&d.Data); err != nil {
			log.Println(err)
			return DataObject{}, nil
		}
		return d, nil
	}
}

func SaveShipment(ctx context.Context, cfg Config, dataobject DataObject) (DataObject, error) {
	c := restclient.NewRestClient(
		restclient.WithEndpoint(cfg.Endpoint),
		restclient.WithAccessToken(cfg.AccessToken),
	)
	// Marshal the shipment
	s, err := json.Marshal(dataobject)
	if err != nil {
		return DataObject{}, err
	}

	// This method has no body
	if b, err := c.Post(ctx, s); err != nil {
		log.Println(err)
		return DataObject{}, err
	} else {
		var do DataObject
		json.Unmarshal(b, &do.Data)
		return do, nil
	}
}
