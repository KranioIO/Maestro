package storage

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoCredentials struct {
	user     string
	pass     string
	endpoint string
	database string
}

var once sync.Once
var connection *mongo.Client
var credentials *mongoCredentials

// Init ..
func Init() {
	setMongoCredentials("mongodb://localhost:27017", "root", "rootpassword", "maestrodb")
	getMongoConn()

	if checkMongoDBConnIsOk() {
		log.Println("connection to MongoDB established")
	}
}

func setMongoCredentials(endpoint string, user string, pass string, db string) {
	credentials = &mongoCredentials{user: user, pass: pass, endpoint: endpoint, database: db}
}

func getMongoConn() *mongo.Client {
	if credentials == nil {
		return nil
	}

	once.Do(func() {
		clientOptions := options.Client().ApplyURI(credentials.endpoint).SetAuth(options.Credential{
			AuthMechanism: "SCRAM-SHA-1", Username: credentials.user, Password: credentials.pass,
		})

		// ctx, _ = context.WithTimeout(context.Background(), 15*time.Second)
		client, err := mongo.Connect(context.TODO(), clientOptions)

		if err != nil {
			log.Println(err)
		} else {
			connection = client
		}
	})

	return connection
}

func checkMongoDBConnIsOk() bool {
	conn := getMongoConn()
	err := conn.Ping(context.TODO(), nil)

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

// InsertData ..
func InsertData(collectionName string, document interface{}) error {
	if !checkMongoDBConnIsOk() {
		return errors.New("Database ERROR no connection found")
	}

	conn := getMongoConn()
	collection := conn.Database(credentials.database).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	insertResult, err := collection.InsertOne(ctx, document)

	if err != nil {
		log.Println(err)
	}

	log.Println("document inserted: ", insertResult.InsertedID)

	return nil
}
