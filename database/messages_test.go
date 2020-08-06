package database

import (
	"context"
	"testing"
)

func TestListClientMessages(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	db, err := Connect("user=postgres dbname=testing password=cinnamon sslmode=disable")
	if err != nil {
		panic(err)
	}

	user := &User{
		Username: "kshirlandm",
		id:       23,
		Token:    <-db.Tokens,
	}

	if len(db.ListClientThreads(ctx, user)) < 1 {
		t.Error("unable to retrieve message threads for client")
	}

	cancel()
}

func TestListThread(t *testing.T) {
	// doctorid 7, clientid 410
	ctx, cancel := context.WithCancel(context.Background())
	db, err := Connect("user=postgres dbname=testing password=cinnamon sslmode=disable")
	if err != nil {
		panic(err)
	}

	user := &User{
		id:    23,
		Token: <-db.Tokens,
	}

	doctor := &Doctor{
		Username: "ppattrickl",
	}

	if len(db.ListThread(ctx, user, doctor)) < 1 {
		t.Error("unable to retrieve message threads for client")
	}

	cancel()
}

func TestSendMessage(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	db, err := Connect("user=postgres dbname=testing password=cinnamon sslmode=disable")
	if err != nil {
		panic(err)
	}

	user := &User{
		id:    23,
		Token: <-db.Tokens,
	}

	doctor := &Doctor{
		ID:       22,
		Username: "ppattrickl",
	}

	message := &Message{
		Content: "Lorum ipsum dolor sit amet",
	}

	if e := db.SendMessage(ctx, user, doctor, message); e != nil {
		t.Error(e)
	}

	cancel()
}
