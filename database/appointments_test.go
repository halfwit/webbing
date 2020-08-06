package database

import (
	"context"
	"testing"
	"time"
)

func TestAddAppointment(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	db, err := Connect("user=postgres dbname=testing password=cinnamon sslmode=disable")
	if err != nil {
		panic(err)
	}

	doctor := &Doctor{
		Username: "johndoe",
	}

	req := &ApptRequest{
		Block:  time.Now().Round(time.Minute * 15),
		Doctor: doctor,
		Token:  <-db.Tokens,
		ID:     42,
	}

	if e := db.AddAppointment(ctx, req); e != nil {
		t.Error(e)
	}

	cancel()
}

// John Doe works out of Antarctica
func TestBookAppointment(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	db, err := Connect("user=postgres dbname=testing password=cinnamon sslmode=disable")
	if err != nil {
		panic(err)
	}

	doctor := &Doctor{
		Username: "johndoe",
	}

	user := &User{
		Username: "ebreukelman1",
	}

	req := &ApptRequest{
		Doctor: doctor,
		User:   user,
		Token:  <-db.Tokens,
		ID:     42,
	}

	if e := db.BookAppointment(ctx, req); e != nil {
		t.Error(e)
	}

	cancel()
}

func TestListAvailable(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	db, err := Connect("user=postgres dbname=testing password=cinnamon sslmode=disable")
	if err != nil {
		panic(err)
	}

	doctor := &Doctor{
		Username: "spassman1",
	}

	req := &ApptRequest{
		Doctor: doctor,
		Token:  <-db.Tokens,
		ID:     42,
	}

	appts := db.ListAvailable(ctx, req)
	if len(appts) < 1 {
		t.Error("Unable to find test appointments")
	}

	cancel()
}

func TestListClientBookings(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	db, err := Connect("user=postgres dbname=testing password=cinnamon sslmode=disable")
	if err != nil {
		panic(err)
	}

	user := &User{
		Username: "ebreukelman1",
		Token:    <-db.Tokens,
	}

	appts := db.ListClientBookings(ctx, user)
	if len(appts) < 1 {
		t.Error("Unable to find test bookings")
	}

	cancel()
}
