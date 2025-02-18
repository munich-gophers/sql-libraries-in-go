package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type User struct {
	ID        int64     `bun:",pk,autoincrement"`
	Name      string    `bun:"name,notnull"`
	Birthday  time.Time `bun:"birthday,notnull"`
	CreatedAt time.Time `bun:"created_at,notnull"`
	UpdatedAt time.Time `bun:"updated_at,notnull"`
	DeletedAt time.Time `bun:"deleted_at,notnull"`
}

func main() {
	// Connect to PostgreSQL
	conn := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(os.Getenv("POSTGRES_DSN"))))
	db := bun.NewDB(conn, pgdialect.New())

	// Create table
	ctx := context.Background()
	// err := db.NewCreateTable().Model((*User)(nil)).Exec(ctx)

	// Insert user
	// user := &User{Name: "John Doe", Email: "john@example.com"}
	// _, err := db.NewInsert().Model(user).Exec(ctx)
	// if err != nil {
	// 	log.Fatalf("Failed to insert user: %v", err)
	// }

	// Fetch user
	var users []User
	err := db.NewSelect().Model(&users).Scan(ctx)
	if err != nil {
		log.Fatalf("Failed to fetch user: %v", err)
	}
	fmt.Printf("Query Results (%d)\n", len(users))
	for _, entry := range users {
		spew.Dump(entry)
	}
}
