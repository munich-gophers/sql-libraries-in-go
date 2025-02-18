package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Age      int64 `gorm:"default:18"`
	Birthday time.Time
}

func main() {
	dsn := os.Getenv("POSTGRES_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
	}
	users := []*User{
		{Name: "Jinzhu", Age: 18, Birthday: time.Now()},
		{Name: "Jackson", Age: 19, Birthday: time.Now()},
	}

	db.Create(users)

	// Example query using GORM
	var result []User
	status := db.Find(&result)

	if status.Error != nil {
		log.Fatalf("coulnd't retrieve database entries: %s\n", status.Error)
	}

	fmt.Printf("Query Results (%d)\n", len(result))
	for _, entry := range result {
		spew.Dump(entry)
	}
}
