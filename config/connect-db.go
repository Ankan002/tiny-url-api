package config

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func ConnectToDB() {
	dbUri := os.Getenv("DB_URI")
	err := mgm.SetDefaultConfig(nil, "PrimaryDatabase", options.Client().ApplyURI(dbUri))

	if err != nil {
		log.Fatal("DB Connection Failed....")
	}

	log.Println("DB Connected....")
}
