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
	podcastResult, err := podcastsCollection.InsertOne(ctx, bson.D{
		{"title", "The Polyglot Developer Podcast"},
		{"author", "Nic Raboy"},
		{"tags", bson.A{"development", "programming", "coding"}},
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(podcastResult.InsertedID)

	//Insert many BSON documents
	episodeResult, err := episodesCollection.InsertMany(ctx, []interface{}{
		bson.D{
			{"podcast", podcastResult.InsertedID},
			{"title", "Episode #01"},
			{"description", "This is the first episode"},
			{"duration", 25},
		},
		bson.D{
			{"podcast", podcastResult.InsertedID},
			{"title", "Episode #02"},
			{"description", "This is the second episode"},
			{"duration", 32},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodeResult.InsertedIDs)
}
