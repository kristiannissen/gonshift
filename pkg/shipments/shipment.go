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

type Address struct {
	Kind        int    `json:"Kind"`
	Name1       string `json:"Name1"`
	Street1     string `json:"Street1"`
	PostCode    string `json:"PostCode"`
	City        string `json:"City"`
	Phone       string `json:"Phone"`
	Mobile      string `json:"Mobile"`
	Email       string `json:"Email"`
	Attention   string `json:"Attention"`
	CountryCode string `json:"CountryCode"`
	Residential bool   `json:"Residential"`
}

type Reference struct {
	Kind  string `json:"Kind"`
	Value string `json:"Value"`
}

type CSEvent struct {
	Kind     int    `json:"Kind"`
	Date     string `json:"Date"`
	User     string `json:"User"`
	EventTag string `json:"EventTag"`
}

type Pkg struct {
	PkgCSID  int       `json:"PkgCSID"`
	ItemNo   int       `json:"ItemNo"`
	PkgTag   string    `json:"PkgTag"`
	PkgNo    string    `json:"PkgNo"`
	CSEvents []CSEvent `json:"CSEvents"`
}

type Line struct {
	LineCSID   int         `json:"LineCSID"`
	LineTag    string      `json:"LineTag"`
	LineWeight int         `json:"LineWeight"`
	PkgWeight  int         `json:"PkgWeight"`
	Pkgs       []Pkg       `json:"Pkgs"`
	References []Reference `json:"References"`
}

type ShpDocument struct {
	ShpCSID      int    `json:"ShpCSID"`
	DocumentID   int    `json:"DocumentID"`
	DocumentCSID int    `json:"DocumentCSID"`
	Copies       int    `json:"Copies"`
	DocumentName string `json:"DocumentName"`
	PrintedDt    string `json:"PrintedDt"`
	DocumentType string `json:"DocumentType"`
}

type Shipment struct {
	ShpCSID                string        `json:"ShpCSID"`
	ShpTag                 string        `json:"ShpTag"`
	InstallationID         string        `json:"InstallationID"`
	PhysicalInstallationID string        `json:"PhysicalInstallationID"`
	Kind                   int           `json:"Kind"`
	ShpNo                  string        `json:"ShpNo"`
	OrderNo                string        `json:"OrderNo"`
	PickupDt               string        `json:"PickupDt"`
	LabelPrintDt           string        `json:"LabelPrintDt"`
	SubmitDt               string        `json:"SubmitDt"`
	TransmitDt             string        `json:"TransmitDt"`
	Weight                 int           `json:"Weight"`
	ActorCSID              int           `json:"ActorCSID"`
	Temperature            int           `json:"Temperature"`
	CarriagePayer          int           `json:"CarriagePayer"`
	CarrierConceptID       int           `json:"CarrierConceptID"`
	CarrierCSID            int           `json:"CarrierCSID"`
	SubcarrierConceptID    int           `json:"SubcarrierConceptID"`
	SubcarrierCSID         int           `json:"SubcarrierCSID"`
	ProdConceptID          int           `json:"ProdConceptID"`
	ProdCSID               int           `json:"ProdCSID"`
	StackCSID              int           `json:"StackCSID"`
	PayerAccountAtCarrier  string        `json:"PayerAccountAtCarrier"`
	SenderAccountAtCarrier string        `json:"SenderAccountAtCarrier"`
	Addresses              []Address     `json:"Addresses"`
	References             []Reference   `json:"References"`
	CSEvents               []CSEvent     `json:"CSEvents"`
	Services               []int         `json:"Services"`
	ShpDocuments           []ShpDocument `json:"ShpDocuments"`
	CorrelationID          string        `json:"CorrelationID"`
}

func GetShipment(ctx context.Context, cfg Config) (Shipment, error) {
	c := restclient.NewRestClient(
		restclient.WithEndpoint(cfg.Endpoint),
		restclient.WithAccessToken(cfg.AccessToken),
	)
	// This method has no body
	if b, err := c.Get(ctx, []byte(``)); err != nil {
		log.Fatal(err)
		return Shipment{}, err
	} else {
		s := Shipment{}
		json.Unmarshal(b, &s)
		return s, nil
	}
}
