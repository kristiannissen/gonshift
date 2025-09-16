package shipments

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"testing"
)

func TestPostShipment(t *testing.T) {
	//
	shipment := Shipment{
		Data: map[string]any{
			"Kind":    1,
			"ActorId": 2,
			"Name":    "Kitty",
		},
		Options: map[string]any{
			"Labels": "PDF",
		},
	}

	actorId, _ := strconv.Atoi(os.Getenv("ACTOR_ID"))
	blob, e := PostShipments(actorId, &shipment)
	if e != nil {
		log.Fatalf("e %s", e)
	}
	want := "Kitty"

	var s Shipment

	if err := json.Unmarshal([]byte(blob), &s); err != nil {
		log.Fatalf("Json %T", err)
	}

	if want != s.Data["Name"] {
		t.Errorf("Got %s, want %s", shipment, want)
	}
}
