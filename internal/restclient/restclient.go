package restclient

//
//
//
//
import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
)

func init() {
}

// TODO: Add boolean for debug
type RestClient struct {
	endpoint string
	token    string
	payload  []byte
}

type Option func(*RestClient)

func NewRestClient(opts ...Option) *RestClient {
	r := &RestClient{
		endpoint: "",
		token:    "",
		payload:  []byte("{}"),
	}

	for _, opt := range opts {
		opt(r)
	}

	return r
}

func WithEndPoint(endpoint string) Option {
	return func(r *RestClient) {
		r.endpoint = endpoint
	}
}

func WithToken(token string) Option {
	return func(r *RestClient) {
		r.token = token
	}
}

func WithPayload(payload []byte) Option {
	return func(r *RestClient) {
		r.payload = payload
	}
}

// TODO: add context
func (r *RestClient) Post() ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, r.endpoint, bytes.NewReader(r.payload))
	if err != nil {
		log.Fatal(err)
		return []byte{}, err
	}

	err = nil

	// Add headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.token))

	err = nil

	d, err := httputil.DumpRequest(req, true)
	if err != nil {
		log.Fatal(err)
		return []byte{}, err
	}

	log.Println(string(d))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return []byte{}, err
	}
	defer res.Body.Close()

	err = nil

	if res.StatusCode != 200 {
		log.Fatal(res.Status)
		return []byte{}, errors.New(res.Status)
	}

	b, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
		return []byte{}, err
	}

	return b, nil
}
