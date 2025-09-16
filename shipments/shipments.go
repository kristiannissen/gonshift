/*
 */
package shipments

import (
	"encoding/json"
	"log"
)

type Shipment struct {
	Data    map[string]interface{} `json:"data,omitempty"`
	Options map[string]interface{} `json:"options,omitempty"`
}

func PostShipments(actorId int, shipment *Shipment) (string, error) {
	b, err := json.MarshalIndent(shipment, "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	return string(b), err
}
