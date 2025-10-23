package models

type DataObject struct {
	Data    map[string]interface{} `json:"data"`
	Options map[string]interface{} `json:"options,omitempty"`
}
