package models

import (
	"context"
	"sync"

	"cloud.google.com/go/datastore"
	"google.golang.org/appengine"
)

type Note struct {
	Address string
	Text    string
}

var localNotes = struct {
	sync.RWMutex
	notes map[string]string
}{
	notes: make(map[string]string),
}

func GetNote(ctx context.Context, address string) (string, error) {
	if !appengine.IsAppEngine() {
		localNotes.RLock()
		text, ok := localNotes.notes[address]
		localNotes.RUnlock()
		if !ok {
			return "", datastore.ErrNoSuchEntity
		}
		return text, nil
	}

	projectID := appengineAppID(ctx)

	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		return "", err
	}
	defer client.Close()

	noteKey := datastore.NameKey("awacha-store", address, nil)

	var note Note
	if err := client.Get(ctx, noteKey, &note); err != nil {
		return "", err
	}

	return note.Text, nil
}

func PutNote(ctx context.Context, note Note) (*datastore.Key, error) {
	noteKey := datastore.NameKey("awacha-store", note.Address, nil)
	if !appengine.IsAppEngine() {
		localNotes.Lock()
		localNotes.notes[note.Address] = note.Text
		localNotes.Unlock()
		return noteKey, nil
	}

	projectID := appengineAppID(ctx)

	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	key, err := client.Put(ctx, noteKey, &note)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func DeleteNote(ctx context.Context, address string) error {
	if !appengine.IsAppEngine() {
		localNotes.Lock()
		delete(localNotes.notes, address)
		localNotes.Unlock()
		return nil
	}

	projectID := appengineAppID(ctx)

	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()

	nameKey := datastore.NameKey("awacha-store", address, nil)

	if err := client.Delete(ctx, nameKey); err != nil {
		return err
	}

	return nil
}

func appengineAppID(ctx context.Context) string {
	if appengine.IsAppEngine() {
		return appengine.AppID(ctx)
	}
	return "awacha-com"
}
