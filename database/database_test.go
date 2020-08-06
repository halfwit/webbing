package database

import (
	"testing"
)

func TestTokenTimeout(t *testing.T) {
	db, err := Connect("user=postgres dbname=testing password=cinnamon sslmode=disable")
	if err != nil {
		t.Error(err)
	}

	// Grab a token and make sure it's valid
	db.Close()
}
