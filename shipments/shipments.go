/*

*/
package shipments

type Shipment struct {
  Data struct {} `json:"data"`
  Options struct {
    Labels string `sjon:"Labels"`
  } `json:"options"`
}

func PostShipments(actorId int, shipment *Shipment) (string, error) {
  return "Hello", nil
}
