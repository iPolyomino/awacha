package models

import (
	"context"
	"errors"
	"testing"

	"cloud.google.com/go/datastore"
)

func TestLocalNoteLifecycle(t *testing.T) {
	localNotes.Lock()
	localNotes.notes = make(map[string]string)
	localNotes.Unlock()

	ctx := context.Background()
	note := Note{Address: "example", Text: "first"}

	key, err := PutNote(ctx, note)
	if err != nil {
		t.Fatalf("PutNote() error = %v", err)
	}
	if key.Name != note.Address {
		t.Errorf("PutNote() key name = %q, want %q", key.Name, note.Address)
	}

	text, err := GetNote(ctx, note.Address)
	if err != nil {
		t.Fatalf("GetNote() error = %v", err)
	}
	if text != note.Text {
		t.Errorf("GetNote() = %q, want %q", text, note.Text)
	}

	note.Text = "updated"
	if _, err := PutNote(ctx, note); err != nil {
		t.Fatalf("PutNote() overwrite error = %v", err)
	}
	text, err = GetNote(ctx, note.Address)
	if err != nil {
		t.Fatalf("GetNote() after overwrite error = %v", err)
	}
	if text != note.Text {
		t.Errorf("GetNote() after overwrite = %q, want %q", text, note.Text)
	}

	if err := DeleteNote(ctx, note.Address); err != nil {
		t.Fatalf("DeleteNote() error = %v", err)
	}
	if _, err := GetNote(ctx, note.Address); !errors.Is(err, datastore.ErrNoSuchEntity) {
		t.Errorf("GetNote() after delete error = %v, want %v", err, datastore.ErrNoSuchEntity)
	}
}
