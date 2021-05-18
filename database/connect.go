package database

import (
	"context"
	"log"
	"time"

	"github.com/pauloprescendo/go-rest-api/pauloprescendo/rest-api-go/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.Config("MONGO_URI")))

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)
}
