package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//Struct for storing data

type user struct {
	Name string `json:name`
	Age  int    `json:age`
	City string `json:city`
}

//Get users collection from db() which returns mongo.Client
var userCollection = db().Database("quickstart-mongodb").Collection("users")

func createProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person user
	//Storing in person variable of type user
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		fmt.Print(err)
	}
	insertResult, err := userCollection.InsertOne(context.TODO(), person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted a single document: ", insertResult)
	json.NewEncoder(w).Encode(insertResult.InsertedID)
	//return the mongodb id of geenrated document
}
