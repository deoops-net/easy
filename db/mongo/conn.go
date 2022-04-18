package mongo

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client
var db *mongo.Database
var mongoUri string

type MongoConf struct {
	DB       string
	Host     string
	Port     string
	AuthDB   string
	User     string
	Password string
}

// ObjectID create mongo ObjectId from string or create new one
func ObjectID(id string) primitive.ObjectID {
	if id == "" {
		return primitive.NewObjectID()
	}
	oid, _ := primitive.ObjectIDFromHex(id)
	return oid
}

func InitConn(c MongoConf) {
	if c.User != "" {
		mongoUri = fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", c.User, c.Password, c.Host, c.Port, c.AuthDB)
	} else {
		mongoUri = fmt.Sprintf("mongodb://%s:%s", c.Host, c.Port)
	}

	client = getClient()
	db = client.Database(c.DB)
}

func Table(name string) *mongo.Collection {
	return db.Collection(name)
}

// GetMongoDB ...
func getClient() *mongo.Client {
	if os.Getenv("MONGO_RS_URI") != "" {
		mongoUri = os.Getenv("MONGO_RS_URI")
	}
	fmt.Printf("MONGO CONNECTING TO %s\n", mongoUri)
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	return client
}
