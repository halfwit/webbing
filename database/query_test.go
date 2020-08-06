package database

import (
	"context"
	"testing"
	"time"
)

func TestAuxQuery(t *testing.T) {
	// Ensure token creation and running works as expected
	db, err := Connect("user=postgres dbname=testing password=cinnamon sslmode=disable")
	if err != nil {
		t.Error(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	q := &Query{
		Country: "Afghanistan",
		Query:   AuxEmail,
		Token:   <-db.Tokens,
	}

	for d := range db.RunQuery(ctx, q) {
		if d != nil {
			cancel()
			return
		}
	}

	cancel()
	t.Error("found no results")
}

func TestNormalQuery(t *testing.T) {
	// Ensure token creation and running works as expected
	db, err := Connect("user=postgres dbname=testing password=cinnamon sslmode=disable")
	if err != nil {
		t.Error(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := &Query{
		Country: "Afghanistan",
		Query:   AuxNone,
		Token:   <-db.Tokens,
	}

	for d := range db.RunQuery(ctx, q) {
		if d != nil {
			return
		}
	}

	t.Error("found no results")
}
