package main

// Refactor
// TODO: follow https://www.alexedwards.net/blog/how-to-manage-configuration-settings-in-go-web-applications
import (
	"context"
	"flag"
	"fmt"
	"os"
	// "log"
	// "bufio"
	"gonshift/pkg/authenticate"
	"path/filepath"
)

func main() {
	/*
	 * Create the .env file
	 */
	d, _ := os.Getwd()
	io, err := os.Create(filepath.Join(d, ".env"))
	if err != nil {
		panic(err)
	}
	// Defer file close
	defer io.Close()
	/*
	 * Flags needed:
	 * - client_id string
	 * - client_secret string
	 * - actor_id int
	 */
	var (
		clientId, clientSecret string
		actorId                int
	)

	flag.StringVar(&clientId, "CLIENT_ID", "Uuid", "Client ID")
	flag.StringVar(&clientSecret, "CLIENT_SECRET", "Uuid", "Client Secret")
	flag.IntVar(&actorId, "ACTOR_ID", 0, "Actor ID")
	// Parse user flag input
	flag.Parse()

	// Run the Authenticate method and store the token in .env
	cfg := authenticate.Config{
		ClientId:     clientId,
		ClientSecret: clientSecret,
	}
	if cfg, err := authenticate.Authenticate(context.Background(), "https://account.nshiftportal.com/idp/connect/token", cfg); err != nil {
		panic(err)
	} else {
		if _, err := io.WriteString(fmt.Sprintf("ACCESS_TOKEN='%s'", cfg.Access_Token)); err != nil {
			panic(err)
		}
	}

	// Iterate over the flags that have been set
	flag.Visit(func(f *flag.Flag) {
		// fmt.Printf("Flag %s: %v - %v", f.Name, f.DefValue, f.Value)
		// Use '' because dotenv has an issue with leading $ in strings
		_, e := io.WriteString(fmt.Sprintf("%s='%s'\n", f.Name, f.Value))
		if e != nil {
			panic(e)
		}
	})

	//
	io.Sync()
	//
	fmt.Println(fmt.Sprintf("Environment variables saved in %s", filepath.Join(d, ".env")))
}
