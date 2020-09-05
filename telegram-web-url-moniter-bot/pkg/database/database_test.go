package database

import (
	"testing"
)

func TestOpenOrCreateDatabase(t *testing.T) {
	t.Parallel()
	_, err := OpenOrCreateDatabase("test.db")
	if err != nil {
		t.Fatal(err)
	}
	RemoveDatabase("test.db")
}