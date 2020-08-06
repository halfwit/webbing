package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Database struct {
	Tokens chan string
	Temps  chan string

	// internal
	db    *sql.DB
	valid map[string]time.Time
}

// Open the socket
func Connect(credentials string) (*Database, error) {
	db, err := sql.Open("postgres", credentials)
	if err != nil {
		return nil, err
	}

	// Make sure we can actually ping the thing
	if e := db.Ping(); e != nil {
		return nil, e
	}

	d := &Database{
		valid: make(map[string]time.Time),
		db:    db,
	}

	d.Tokens = createTokens(d)
	d.Temps = createTemps(d)
	return d, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func createTokens(d *Database) chan string {
	tokens := make(chan string)
	go func(tokens chan string) {
		for {
			u, _ := uuid.NewRandom()
			t := u.String()

			tokens <- t
			// After token unblocks, add a timeout to it
			// Plugins are destroyed on-use, the only case
			// where errant tokens would hang around is if
			// a plugin is programmed incorrectly
			d.valid[t] = time.Now().Add(time.Second * 5)
		}
	}(tokens)

	return tokens
}

func createTemps(d *Database) chan string {
	tokens := make(chan string)
	go func(tokens chan string) {
		for {
			u, _ := uuid.NewRandom()
			t := u.String()

			tokens <- t

			// Temp tokens are internally destruct-on-use
			// But in general a signup, or otherwise may take
			// quite a while. We give a forgiving timeout
			d.valid[t] = time.Now().Add(time.Hour)
		}
	}(tokens)

	return tokens
}
