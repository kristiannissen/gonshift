package shipments

import (
	"os"
	"strconv"
	"testing"
)

func TestPostShipment(t *testing.T) {
	//
	shipment := Shipment{}

	actorId, _ := strconv.Atoi(os.Getenv("ACTOR_ID"))
	got, _ := PostShipments(actorId, &shipment)
	want := "Hello"

	if want != got {
		t.Errorf("Got %s, want %s", got, want)
	}
}
