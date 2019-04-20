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

func GetNote(ctx context.Context, address string) (string, error) {
	projectID := appengineAppID(ctx)

	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Could not create datastore client: %v", err)
		return "", err
	}

	noteKey := datastore.NameKey("awacha-store", address, nil)

	var note Note
	if err := client.Get(ctx, noteKey, &note); err != nil {
		return "", err
	}

	return note.Text, nil
}

func PutNote(ctx context.Context, note Note) (*datastore.Key, error) {
	projectID := appengineAppID(ctx)

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

func DeleteNote(ctx context.Context, address string) error {
	projectID := appengineAppID(ctx)

	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Could not create datastore client: %v", err)
		return err
	}

	nameKey := datastore.NameKey("awacha-store", address, nil)

	if err := client.Delete(ctx, nameKey); err != nil {
		log.Fatalf("Could not delete %v: %v", address, err)
		return err
	}

	return nil
}

func appengineAppID(ctx context.Context) string {
	projectID := appengine.AppID(ctx)
	if projectID == "None" {
		projectID = "awacha-com"
	}

	return projectID
}
