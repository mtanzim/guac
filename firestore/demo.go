package firestore

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/mtanzim/guac/wakaApi"
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

func Get(collName string, ctx context.Context, client *firestore.Client) {
	iter := client.Collection(collName).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		log.Println(doc.Data())
	}
}

func Demo() {
	ctx := context.Background()
	client, close := CreateClient(ctx)
	defer close()

	data := wakaApi.TransformData()
	var items []Item
	for k, v := range data {
		item := Item{"mtanzim", k, v}
		items = append(items, item)
	}

	collName := os.Getenv("GOOGLE_WAKA_COLL")
	if collName == "" {
		log.Fatalf("Please setup firestore collection name env var")
	}

	Put(collName, items, ctx, client)
	// Get(collName, ctx, client)

}
