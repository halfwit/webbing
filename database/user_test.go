package database

import (
	"context"
	"testing"
	"time"
)

func TestUpdateUser(t *testing.T) {

}

func TestFindUser(t *testing.T) {
	// Ensure token creation and running works as expected
	db, err := Connect("user=postgres dbname=testing password=cinnamon sslmode=disable")
	if err != nil {
		t.Error(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	u := &UserInfo{
		// Random pull from mock database
		Username: "jcoyishix",
		Password: "cUco9wEq",
		Token:    <-db.Tokens,
	}

	if user := db.FindUser(ctx, u); user != nil {
		return
	}

	t.Error("found no results")
}
