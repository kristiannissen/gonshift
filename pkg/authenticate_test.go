package main

import (
	"os"
	"testing"
)

func TestAuthenticate(t *testing.T) {

	got, _ := Authenticate(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))
	want := ""

	if want >= got {
		t.Errorf("Got %s but want %s", got, want)
	}
}
