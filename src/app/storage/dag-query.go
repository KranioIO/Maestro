package storage

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetDagExecutionsByName ..
func GetDagExecutionsByName() ([]map[string]interface{}, error) {
	conn := getMongoConn()
	collection := conn.Database(credentials.database).Collection("executions")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	findOptions := options.Find()
	findOptions.SetLimit(10)

	cur, err := collection.Find(ctx, bson.D{{}}, findOptions)

	if err != nil {
		return nil, err
	}

	results := make([]map[string]interface{}, 0, 10)

	for cur.Next(ctx) {
		elem := make(map[string]interface{})

		if err := cur.Decode(elem); err != nil {
			return nil, err
		}

		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	cur.Close(ctx)

	return results, nil
}
