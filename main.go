//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/SSHcom/privx-sdk-go/api"
	"github.com/SSHcom/privx-sdk-go/config"
	"github.com/SSHcom/privx-sdk-go/oauth"
)

var commands = map[string]func(client *api.Client){
	"login": cmdLogin,
	"get":   cmdGet,
}

func main() {
	log.SetFlags(0)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr,
			"Usage: %s [options] COMMAND [command options] [ARG]...\n",
			os.Args[0])
		fmt.Fprintf(os.Stderr, "\nOptions:\n")
		flag.PrintDefaults()

		fmt.Fprintf(os.Stderr, "\nCommands:\n")
		for key := range commands {
			fmt.Fprintf(os.Stderr, "  - %s\n", key)
		}
		fmt.Fprintf(os.Stderr,
			"\nType %s COMMAND -h for help about COMMAND\n",
			os.Args[0])
	}

	apiEndpoint := flag.String("api", "", "API endpoint URL")
	configFile := flag.String("config", config.Default(), "configuration file")
	verbose := flag.Bool("v", false, "verbose output")
	flag.Parse()

	config, err := config.Read(*configFile)
	if err != nil {
		log.Fatalf("Failed to read config file '%s': %s", *configFile, err)
	}

	// Command line overrides.
	if len(*apiEndpoint) > 0 {
		config.API.Endpoint = *apiEndpoint
	}

	// Construct API client.
	auth, err := oauth.NewClient(config.Auth, config.API.Endpoint,
		config.API.Certificate.X509, *verbose)
	if err != nil {
		log.Fatal(err)
	}
	client, err := api.NewClient(auth, config.API.Endpoint,
		config.API.Certificate.X509, *verbose)
	if err != nil {
		log.Fatal(err)
	}

	if len(flag.Args()) == 0 {
		fmt.Fprintf(os.Stderr, "No command specified.\n")
		return
	}
	os.Args = flag.Args()
	fn, ok := commands[flag.Arg(0)]
	if !ok {
		fmt.Printf("Unknown command: %s\n", flag.Arg(0))
		os.Exit(1)
	}
	flag.CommandLine = flag.NewFlagSet(
		fmt.Sprintf("privx-secrets %s", os.Args[0]),
		flag.ExitOnError)
	fn(client)
}
