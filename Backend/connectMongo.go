package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func connectMongo(client mongo.Client, txtlines []string, err error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	myDBDatabase := client.Database("myDB")

	docsCollection := myDBDatabase.Collection("docs")

	//Insert many BSON documents
	valResult, err := docsCollection.InsertMany(ctx, []interface{}{
		bson.D{

			{"title", txtlines[0]},
			{"description", txtlines[1]},
			{"duration", 25},
		},
		bson.D{

			{"title", "Episode #02"},
			{"description", "This is the second episode"},
			{"duration", 32},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(valResult.InsertedIDs)
}
