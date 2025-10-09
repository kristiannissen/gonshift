package restclient

// TODO: DRY the fuck out of this!
// TODO: Return error message from body
//
//
import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
)

func init() {
}

type ErrorMessage struct {
	ErrorType string   `json:"ErrorType"`
	Message   string   `json:"Message"`
	Errors    []string `json:"Errors"`
}

type RestClient struct {
	Endpoint    string
	AccessToken string
	Payload     []byte
	Debug       bool
}

type Option func(*RestClient)

func NewRestClient(opts ...Option) *RestClient {
	r := &RestClient{
		Endpoint:    "",
		AccessToken: "",
		Debug:       false,
	}

	for _, opt := range opts {
		opt(r)
	}

	return r
}

func WithEndpoint(endpoint string) Option {
	return func(r *RestClient) {
		r.Endpoint = endpoint
	}
}

func WithAccessToken(token string) Option {
	return func(r *RestClient) {
		r.AccessToken = token
	}
}

func WithDebug(b bool) Option {
	return func(r *RestClient) {
		r.Debug = b
	}
}

func (r *RestClient) Post(ctx context.Context, payload []byte) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, r.Endpoint, bytes.NewBuffer(payload))
	if err != nil {
		log.Println(err)
		return []byte(``), err
	}

	// Create headers and authorization
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("bearer %s", r.AccessToken))

	if r.Debug == true {
		// Debug is enabled
		if rd, err := httputil.DumpRequestOut(req, true); err != nil {
			log.Println(err)
		} else {
			log.Println(string(rd))
		}
	}

	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		log.Println(err)
		// Check ctx error
		if ctx.Err() == context.DeadlineExceeded {
			return []byte(``), ctx.Err()
		}
		return []byte(``), err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Println(res.Status)
		var msg ErrorMessage
		if err := json.NewDecoder(res.Body).Decode(&msg); err != nil {
			log.Println(err)
			return []byte(``), err
		} else {
			// Return the error message from the server
			log.Println(msg.Message)
			return []byte(``), errors.New(msg.Message)
		}
	}

	if b, err := io.ReadAll(res.Body); err != nil {
		log.Println(err)
		return []byte(``), err
	} else {
		return b, nil
	}
}

func (r *RestClient) Get(ctx context.Context, payload []byte) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, r.Endpoint, bytes.NewBuffer(payload))
	if err != nil {
		log.Println(err)
		return []byte(``), err
	}

	// Create headers and authorization
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("bearer %s", r.AccessToken))

	if r.Debug == true {
		// Debug is enabled
		if rd, err := httputil.DumpRequestOut(req, true); err != nil {
			log.Println(err)
		} else {
			log.Println(string(rd))
		}
	}

	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		log.Println(err)
		// Check ctx error
		if ctx.Err() == context.DeadlineExceeded {
			return []byte(``), ctx.Err()
		}
		return []byte(``), err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Println(res.Status)
		var msg ErrorMessage
		if err := json.NewDecoder(res.Body).Decode(&msg); err != nil {
			log.Println(err)
			return []byte(``), err
		} else {
			log.Println(msg.Message)
			// Return the error message from the server
			return []byte(``), errors.New(msg.Message)
		}
	}

	if b, err := io.ReadAll(res.Body); err != nil {
		log.Println(err)
		return []byte(``), err
	} else {
		return b, nil
	}
}
