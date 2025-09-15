package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var client_id, client_secret string
	flag.StringVar(&client_id, "client_id", "", "The client id")
	flag.StringVar(&client_secret, "client_secret", "", "The client secret")

	flag.Parse()

	// Store flags as environment variables
	os.Setenv("CLIENT_ID", client_id)
	os.Setenv("CLIENT_SECRET", client_secret)

	fmt.Println("Environment variables set, you can now run the authentication method")
}
