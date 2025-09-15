/*
 */
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type token struct {
	Access_Token string
}

type statusError struct {
	m string
}

func (e *statusError) Error() string {
	return e.m
}

func Authenticate(clientId string, clientSecret string) (string, error) {
	endpoint := "https://account.nshiftportal.com/idp/connect/token"

	resp, err := http.PostForm(endpoint, url.Values{
		"CLIENT_ID":     {clientId},
		"CLIENT_SECRET": {clientSecret},
		"grant_type":    {"client_credentials"},
	})

	if err != nil {
		log.Fatalf("Error %s", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Status code: %d, %s", resp.StatusCode, resp.Status)
		return "", &statusError{m: fmt.Sprintf("%s", resp.Status)}
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	var t token
	json.Unmarshal(data, &t)

	return t.Access_Token, nil
}
