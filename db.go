package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func db() *mongo.Client {

	//Connect to MongoDB
	clientOptions := options.Client().ApplyURI("mongodb+srv://quickstart-mongodb-user:hksY897sT@restaurants-near-me-clu.bczxm.mongodb.net/quickstart-mongodb?retryWrites=true&w=majority")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	//.Further TODO() is used when it’s unclear which Context to use or it is not available
	//Check connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection success to MongoDB")

	return client

}
