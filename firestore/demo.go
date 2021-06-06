package firestore

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type Item struct {
	UserID string
	Date   string
	Data   interface{}
}

func CreateClient(ctx context.Context) (*firestore.Client, func() error) {
	// Sets your Google Cloud Platform project ID.
	projectID := os.Getenv("GOOGLE_PROJECT_ID")
	if projectID == "" {
		log.Fatalf("Please setup google project id env var")
	}

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	// Close client when done with
	return client, client.Close
}

func Put(collName string, items []Item, ctx context.Context, client *firestore.Client) {
	for _, item := range items {
		_, _, err := client.Collection(collName).Add(ctx, item)
		if err != nil {
			log.Fatalf("Failed adding item: %v", err)
		}
	}
}

func Get(collName string, ctx context.Context, client *firestore.Client, start, end string) []Item {

	var rvItems []Item
	iter := client.Collection(collName).Where("Date", ">=", start).Where("Date", "<=", end).Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		item := Item{}
		doc.DataTo(&item)
		rvItems = append(rvItems, item)
	}
	log.Println(rvItems)
	return rvItems

}

func PutData(data map[string]interface{}) {
	collName := os.Getenv("GOOGLE_WAKA_COLL")

	ctx := context.Background()
	client, close := CreateClient(ctx)
	defer close()

	// data := wakaApi.TransformData()
	var items []Item
	for k, v := range data {
		item := Item{"mtanzim", k, v}
		items = append(items, item)
	}
	Put(collName, items, ctx, client)

}

func GetData(start, end string) []Item {
	ctx := context.Background()
	client, close := CreateClient(ctx)
	defer close()

	collName := os.Getenv("GOOGLE_WAKA_COLL")
	if collName == "" {
		log.Fatalf("Please setup firestore collection name env var")
	}

	rv := Get(collName, ctx, client, start, end)
	return rv

}
