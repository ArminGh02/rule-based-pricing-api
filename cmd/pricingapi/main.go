package main

import (
	"log"
	"os"
	"pricingapi/pkg/api"
	"pricingapi/pkg/db"
	"pricingapi/pkg/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	uri := os.Getenv("PRICING_MONGODB_URI")
	if uri == "" {
		log.Fatal("PRICING_MONGODB_URI is not set")
	}

	db := db.New(uri)
	api.Instance().Init(db)

	r := router.New()
	r.Logger.Fatal(r.Start(":8080"))
}
