package configs

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func ConnectDB() (client *mongo.Client, ctx context.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//Validate connection
	err = PingDB(client, ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client, ctx
}

func PingDB(client *mongo.Client, ctx context.Context) error {
	err := client.Ping(ctx, nil)
	if err != nil {
		return err
	}
	fmt.Println("Connected to MongoDB")
	return nil

}

func DisconnectDB(client *mongo.Client, ctx context.Context) {
	err := client.Disconnect(ctx)
	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
	}

}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("golangAPI").Collection(collectionName)
	return collection
}
