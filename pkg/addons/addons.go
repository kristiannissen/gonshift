package addons

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"gonshift/internal/restclient"
)

type Config struct {
	ActorId      string
	Access_Token string
}

type Address struct {
	Kind        int    `json:"Kind"`
	Name1       string `json:"Name1"`
	Street1     string `json:"Street1"`
	PostCode    string `json:"PostCode"`
	City        string `json:"City"`
	CountryCode string `json:"CountryCode"`
}

type Data struct {
	Kind          int   `json:"Kind"`
	ProdConceptID int   `json:"ProdConceptID"`
	Services      []int `json:"Services"`
	Addresses     []Address
}

type DataObject struct {
	Data    Data                   `json:"data"`
	Options map[string]interface{} `json:"options"`
}

type Reference struct {
	Kind  int
	Value string
}

type DropPoint struct {
	OriginalID        string
	RoutingCode       string
	Depot             string
	Name1             string
	Name2             string
	Street1           string
	Street2           string
	PostCode          string
	City              string
	State             string
	Region            string
	County            string
	District          string
	Province          string
	CountryCode       string
	Contact           string
	Phone             string
	Fax               string
	Email             string
	Latitude          float64
	Longitude         float64
	Distance          float32
	Type              string
	DropPointImageUrl string
	DropPointType     string
	References        []Reference
}

// Will post to {{URL}}/ShipServer/{{ID}}/DropPoints
func GetDropPoints(ctx context.Context, endpoint string, cfg Config, payload DataObject) ([]DropPoint, error) {
	endpoint = fmt.Sprintf(endpoint, cfg.ActorId)
	p, _ := json.Marshal(payload)

	c := restclient.NewRestClient(
		restclient.WithEndpoint(endpoint),
		restclient.WithAccessToken(cfg.Access_Token),
	)

	if b, err := c.Post(ctx, p); err != nil {
		log.Fatal(err)
		return []DropPoint{}, err
	} else {
		// Unmarshal and return DropPoints
		d := []DropPoint{}
		json.Unmarshal(b, &d)
		return d, nil
	}
}
