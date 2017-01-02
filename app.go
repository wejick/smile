package main

import (
	"log"

	"github.com/wejick/smile/src/trackingSchema"
	"github.com/wejick/tego/config"
)

func main() {
	err := config.LoadConfigFromFile("./files/etc/smile/config.json")
	if err != nil {
		log.Fatal(err)
	}
	err = trackingSchema.Init()
	if err != nil {
		log.Fatal(err)
	}

	schema, err := trackingSchema.GetSchema("T000001")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(*schema)
}
