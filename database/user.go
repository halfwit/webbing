package database

import (
	"context"
	"errors"
	"time"
)

// Access - Who can access the data
type Access uint8

const (
	GuestAuth Access = 1 << iota
	PatientAuth
	DoctorAuth
)

// User - Any registered user on the site
type User struct {
	Username string
	Country  string
	Token    string
	access   Access
	id       int
}

// UserInfo - request for a user
type UserInfo struct {
	// Must be valid
	Token string
	//Hash string
	Password string
	Username string
}

const createUser = `
INSERT INTO users(name, email, country, password)
VALUES($1, $2, $3, $4)
`

// CreateUser is called from email/signup only
func (d *Database) CreateUser(ctx context.Context, u *User, email, password string) error {
	defer delete(d.valid, u.Token)

	// validate token
	t, ok := d.valid[u.Token]
	if !ok || t.Before(time.Now()) || u.access != PatientAuth {
		return errors.New("invalid/expired access token")
	}

	_, err := d.db.ExecContext(ctx, createUser, u.Username, email, u.Country, password)
	return err
}

const updateUser = `
UPDATE users
SET Country = $2, Username = $3
WHERE users.id = $1`

func (d *Database) UserUpdate(ctx context.Context, u *User) chan struct{} {
	defer delete(d.valid, u.Token)

	// validate token
	t, ok := d.valid[u.Token]
	if !ok || t.Before(time.Now()) || u.access != PatientAuth {
		return make(chan struct{})
	}

	d.db.ExecContext(ctx, updateUser, u.id, u.Country, u.Username)
	return make(chan struct{})
}

const queryUser = `
SELECT users.id, users.name, users.country
FROM users WHERE users.name = $1 AND users.password = $2`

//FROM users WHERE users.name = $1 AND users.hash = $2`

// Attempt to find a user (Most common case by an order of magnitude)
// if we fail, attempt to find a doctor; then fail outright
func (d *Database) FindUser(ctx context.Context, u *UserInfo) *User {
	user := User{
		access:   GuestAuth,
		Username: "none",
		Country:  "none",
	}

	defer delete(d.valid, u.Token)

	// validate token
	t, ok := d.valid[u.Token]
	if !ok || t.Before(time.Now()) {
		return &user
	}

	rows, err := d.db.QueryContext(ctx, queryUser, u.Username, u.Password)
	if err != nil {
		return d.FindDoctor(ctx, u)
	}

	for rows.Next() {
		if e := rows.Scan(&user.id, &user.Username, &user.Country); e != nil {
			break
		}

		// Only set Access on a successful lookup
		user.access = PatientAuth
		break
	}

	return &user
}

const queryDoctor = `
SELECT doctors.nickname, doctors.country
FROM doctors WHERE doctors.nickname = $1 AND doctors.password = $2`

//FROM doctors WHERE doctors.hash = $1`

func (d *Database) FindDoctor(ctx context.Context, u *UserInfo) *User {
	user := User{
		access:   GuestAuth,
		Username: "none",
		Country:  "none",
	}

	rows, err := d.db.QueryContext(ctx, queryUser, u.Username, u.Password)
	if err != nil {
		return &user
	}

	for rows.Next() {
		if e := rows.Scan(&user.Username, &user.Country); e != nil {
			break
		}

		// Only set Access on a successful lookup
		user.access = DoctorAuth
		break
	}

	return &user
}

// Retrieve the access for the given client
func (u *User) Access() Access {
	return u.access
}
