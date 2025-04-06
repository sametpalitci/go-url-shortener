package db

import (
	"context"
	"log"
	"time"

	"url-shortener/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var urlCollection *mongo.Collection

type URL struct {
	ShortURL    string    `bson:"short_url"`
	OriginalURL string    `bson:"original_url"`
	CreatedAt   time.Time `bson:"created_at"`
}

func ConnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURI := utils.GetEnvVar("MONGODB_URI", "mongodb://localhost:27017")
	dbName := utils.GetEnvVar("MONGODB_DB", "urlshortener")
	collectionName := utils.GetEnvVar("MONGODB_COLLECTION", "urls")

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	urlCollection = client.Database(dbName).Collection(collectionName)
	log.Printf("MongoDB | Connected Database: %s, Collection: %s", dbName, collectionName)
}

func SaveURL(url URL) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := urlCollection.InsertOne(ctx, url)
	return err
}

func GetURL(shortURL string) (*URL, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var url URL
	err := urlCollection.FindOne(ctx, bson.M{"short_url": shortURL}).Decode(&url)
	if err != nil {
		return nil, err
	}
	return &url, nil
}

func CloseDB() {
	if client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}
}
