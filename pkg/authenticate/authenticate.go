/*
 * Using the client secret and id, the Authenticate method
 * will post the information and in retur get an access token,
 * this method should probably be run every Nth minute to refresh
 * the access token since it expires every hour.
 *
 * https://account.nshiftportal.com/idp/connect/token
 */
package authenticate

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Config struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	AccessToken  string `json:"access_token"`
	Expires_In   int    `json:"expires_id"`
}

func Authenticate(ctx context.Context, endpoint string, cfg Config) (Config, error) {
	// Create the values
	v := url.Values{
		"CLIENT_ID":     {cfg.ClientId},
		"CLIENT_SECRET": {cfg.ClientSecret},
		"grant_type":    {"client_credentials"},
	}
	// TODO: Could all be moved to a shared feature in the internal folder
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewBuffer([]byte(v.Encode())))
	if err != nil {
		log.Println(err)
		return cfg, fmt.Errorf("request failed %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		// Check context error
		if ctx.Err() == context.DeadlineExceeded {
			return Config{}, fmt.Errorf("the request timed out %w", ctx.Err())
		}
		return Config{}, fmt.Errorf("client failed %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println(resp.StatusCode)
		return cfg, fmt.Errorf("error %s", resp.Status)
	}

	if b, err := io.ReadAll(resp.Body); err != nil {
		// Something went wrong
		log.Println(err)
		return Config{}, err
	} else {
		//
		json.Unmarshal(b, &cfg)

		return cfg, nil
	}
}
