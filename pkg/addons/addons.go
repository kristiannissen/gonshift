package addons

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"gonshift/internal/restclient"
	"gonshift/pkg/models"
)

// GetDropPoints returns drop points (service points) for the given product that supports it.
// Example usage:
//
//	d := GetDropPoints(context.Context, &models.Config{}, &models.DataObject{
//	  "data": "data": {
//	    "ProdConceptID": "INT",
//	    "ResultCount": "INT",
//	    "Addresses": "ARRAY"
//	  }
//	})
//
//	OUTPUT: {
//	   "ProdConceptIDs": [
//	      5318,
//	      5319,
//	      755
//	   ],
//	   "ResultCount": "2",
//	   "Addresses": [
//	      {
//	         "Kind": 1,
//	         "CountryCode": "FR",
//	         "PostCode": "76000",
//	         "Street1": "84, rue de Epeule"
//	      }
//	   ]
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
