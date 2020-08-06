package database

import (
	"context"
	"time"
)

const updateEmail = `
UPDATE users
SET password = $1
WHERE users.email = $2;`

//SET hash = $1
//WHERE users.email = $2;``

type Reset struct {
	Email string
	//Hash string
	Password string
	Token    string
}

func (d *Database) ResetPassword(ctx context.Context, r *Reset) {
	defer delete(d.valid, r.Token)

	// validate token
	t, ok := d.valid[r.Token]
	if !ok || t.Before(time.Now()) {
		return
	}

	d.db.ExecContext(ctx, updateEmail, r.Password, r.Email)
	//d.db.ExecContext(ctx, updateEmail, r.Hash, r.Email)
}
