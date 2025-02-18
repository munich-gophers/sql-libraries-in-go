package main

import (
	"context"
	"ent-demo/ent"
	"fmt"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres", os.Getenv("POSTGRES_DSN"))
	// Initialize the client (e.g., client from ent.Schema)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Example query
	users, err := client.User.Query().All(context.Background())
	if err != nil {
		log.Fatalf("couldn't fetch from database %s\n", err)
	}

	fmt.Printf("Query Results (%d)\n", len(users))
	for _, entry := range users {
		spew.Dump(entry)
	}
}
