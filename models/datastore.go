package models

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
	"google.golang.org/appengine"
)

type Note struct {
	Address string
	Text    string
}

func FetchData(ctx context.Context, address string) (string, error) {
	projectID := appengine.AppID(ctx)
	if projectID == "None" {
		projectID = "awacha-com"
	}

	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Could not create datastore client: %v", err)
		return "", err
	}

	noteKey := datastore.NameKey("awacha-store", address, nil)

	var note Note
	err = client.Get(ctx, noteKey, &note)
	if err != nil {
		return "", err
	}

	return note.Text, nil
}

func StoreingData(ctx context.Context, note Note) (*datastore.Key, error) {
	projectID := appengine.AppID(ctx)
	if projectID == "None" {
		projectID = "awacha-com"
	}

	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Could not create datastore client: %v", err)
		return nil, err
	}

	noteKey := datastore.NameKey("awacha-store", note.Address, nil)

	key, err := client.Put(ctx, noteKey, &note)
	if err != nil {
		log.Fatalf("Could not save %v: %v", note, err)
		return nil, err
	}

	return key, nil
}
