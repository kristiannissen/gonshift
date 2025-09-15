package main

import (
	"os"
	"testing"
)

func TestAuthenticate(t *testing.T) {

	got, err := Authenticate(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))
	want := ""

	if err != nil {
		t.Errorf("Got %s", err)
	}

	if want >= got {
		t.Errorf("Got %s but want %s", got, want)
	}
}
