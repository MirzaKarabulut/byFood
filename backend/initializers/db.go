package initializers

import (
	"context"
	"log"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
    // Get environment variables
    databaseURI := os.Getenv("DATABASE")
    databasePassword := os.Getenv("DATABASE_PASSWORD")

    // Replace <password> placeholder with actual password
    databaseURI = strings.Replace(databaseURI, "<password>", databasePassword, 1)

    // Set client options
    clientOptions := options.Client().ApplyURI(databaseURI)

    // Connect to MongoDB
    client, err := mongo.Connect(context.TODO(), clientOptions)

    if err != nil {
        log.Fatal(err)
    }

    // Check the connection
    err = client.Ping(context.TODO(), nil)

    if err != nil {
        log.Fatal(err)
    }

    return client
}