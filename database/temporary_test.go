package database

import (
	"context"
	"testing"
	"time"
)

func TestCreateTempEntry(t *testing.T) {
	// Ensure token creation and running works as expected
	db, err := Connect("user=postgres dbname=testing password=cinnamon sslmode=disable")
	if err != nil {
		t.Error(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	user := &User{
		Token:    <-db.Tokens,
		Username: "johndoe123_test",
		Country:  "canada",
	}

	if token := db.CreateTempEntry(ctx, user, "foo@bar.com", "hunter2"); token == "BADTOKEN" {
		t.Error("was unable to create a temp entry")
	}

}

func TestForwardToken(t *testing.T) {
	db, err := Connect("user=postgres dbname=testing password=cinnamon sslmode=disable")
	if err != nil {
		t.Error(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if token := db.ForwardToken(ctx, <-db.Temps); token == "BADTOKEN" {
		t.Error("was unable to create new token")
	}
}
