// Copyright 2018 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The newrepo command utilizes go-github as a cli tool for
// creating new repositories. It takes an auth token as
// an environment variable and creates the new repo under
// the account affiliated with that token.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v25/github"
	"golang.org/x/oauth2"
)

var (
	name        = flag.String("name", "", "Name of repo to create in authenticated user's GitHub account.")
	description = flag.String("description", "", "Description of created repo.")
	private     = flag.Bool("private", false, "Will created repo be private.")
)

func main() {
	// parse command-line flags
	flag.Parse()

	// get auth token from environment variables
	token := os.Getenv(
		"GITHUB_TOKEN",
	)

	// panic if token not present
	if token == "" {
		log.Fatal("Unauthorized: No token present")
	}

	// panic if repo name is blank
	if *name == "" {
		log.Fatal("No name: New repos must be given a name")
	}

	// init context
	ctx := context.Background()

	// init new Oauth token
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: token,
		},
	)

	// init new OAuth client
	tokenClient := oauth2.NewClient(
		ctx,
		tokenSource,
	)

	// init new authenticated github client
	client := github.NewClient(tokenClient)

	// init new instance of repository struct
	r := &github.Repository{
		Name:        name,
		Private:     private,
		Description: description,
	}

	// create new repository with authenticated github client
	repo, _, err := client.Repositories.Create(
		ctx,
		"",
		r,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Successfully created new repo: %v\n", repo.GetName())
}
