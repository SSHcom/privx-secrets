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

	"github.com/SSHcom/privx-secrets/api"
)

func cmdLogin(client *api.Client) {
	flag.Parse()

	_, err := client.Auth.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ok")
}
