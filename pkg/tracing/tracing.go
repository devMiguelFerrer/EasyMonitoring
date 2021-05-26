package tracing

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Tracing struct {
	Host           string
	Port           int
	DBName         string
	CollectionName string
	client         *mongo.Client
	collection     *mongo.Collection
}

func (t *Tracing) Connect() {
	var err error
	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d/%s", t.Host, t.Port, t.DBName))
	t.client, err = mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		panic(err)
	}
	t.collection = t.client.Database(t.DBName).Collection(t.CollectionName)
}

func (t *Tracing) Save(i interface{}) {
	insertResult, err := t.collection.InsertOne(context.TODO(), i)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted: ", insertResult.InsertedID)
}
