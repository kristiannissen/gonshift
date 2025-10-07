package restclient

import (
	"testing"
)

func TestPost(t *testing.T) {
	c := NewRestClient(
		WithEndPoint("https://api.restful-api.dev/objects"),
		WithPayload([]byte("{}")),
	)
	_, err := c.Post()

	if err != nil {
		t.Fatal(err)
	}
}
