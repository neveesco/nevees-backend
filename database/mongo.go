package database

import (
	"context"
	"log"

	"github.com/MahdiRazaqi/nevees-backend/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB connection
var MongoDB *mongo.Database

// ConnectMongo Connect to MongoDB
func ConnectMongo() {
	clientOptions := options.Client().ApplyURI(config.CFG.Mongo.Host)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}

	MongoDB = client.Database(config.CFG.Mongo.DB)
}

// ConvertToBson convert interface to bson object
func ConvertToBson(d interface{}) bson.M {
	val, _ := bson.Marshal(d)
	data := new(bson.M)
	bson.Unmarshal(val, data)
	return *data
}
