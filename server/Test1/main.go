package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Retrieve documents in MongoDB

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://neth98:neth123@test.dq1l7.mongodb.net/quickstart?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	quickstartDatabase := client.Database("quickstart")
	podcastsCollection := quickstartDatabase.Collection("podcasts")
	episodesCollection := quickstartDatabase.Collection("episodes")

	cursor, err := episodesCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	//Two ways of loading documents

	// Method 01 - All from a cursor into a slice of our object bson.M

	/*
		var episodes []bson.M
		 if err = cursor.All(ctx, &episodes); err != nil {
		 	log.Fatal(err)
		 }
		 for _, episode := range episodes {
		 	fmt.Println(episode["title"])
		 }
	*/

	//Method 02 - One by one using Next. Uses for large resource set. Overflowing prevented

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var episode bson.M
		if err = cursor.Decode(&episode); err != nil {
			log.Fatal(err)
		}
		//	fmt.Println(episode)
	}

	//Getting Single document
	var podcast bson.M
	if err = podcastsCollection.FindOne(ctx, bson.M{}).Decode(&podcast); err != nil {
		log.Fatal(err)
	}
	//fmt.Println(podcast)

	//Querying Documents from a Collection with a Filter

	filterCursor, err := episodesCollection.Find(ctx, bson.M{"duration": 25}) //findone - more than can come
	if err != nil {
		log.Fatal(err)
	}
	var episodesFiltered []bson.M
	if err = filterCursor.All(ctx, &episodesFiltered); err != nil {
		log.Fatal(err)
	}
	//	fmt.Println(episodesFiltered)

	// Sorting Documents in a Query

	opts := options.Find()
	opts.SetSort(bson.D{{"duration", 1}})                   // -1 -> Descending, 1 -> Ascending
	sortCursor, err := episodesCollection.Find(ctx, bson.D{ //here we use a range query
		{"duration", bson.D{
			{"$gt", 24}, //Greater than 24
		}},
	}, opts)
	if err != nil {
		log.Fatal(err)
	}
	var episodesSorted []bson.M
	if err = sortCursor.All(ctx, &episodesSorted); err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodesSorted)

}

//bson.D -> order of BSON
