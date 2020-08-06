package database

import (
	"context"
	"errors"
	"time"
)

const tempEntry = `
INSERT INTO tmpusers(username, country, email, password, token)
VALUES($1, $2, $3, $4, $5)
`

func (d *Database) CreateTempEntry(ctx context.Context, user *User, email, password string) string {
	defer delete(d.valid, user.Token)

	t, ok := d.valid[user.Token]
	if !ok || t.Before(time.Now()) {
		return "BADTOKEN"
	}

	token := <-d.Temps

	//res, err := d.db.ExecContext(ctx, tempEntry, user.Username, email, password, token)
	//if err != nil { return "BADTOKEN" }
	//if row, e := res.RowsAffected(); e != nil || row < 1 { return "BADTOKEN" }
	if _, e := d.db.ExecContext(ctx, tempEntry, user.Username, user.Country, email, password, token); e != nil {
		return "BADTOKEN"
	}

	return token
}

// If this doesn't work, we'll have to first fetch the user ID
// Then update based on that.  Not a big deal, but still annoying
const updateToken = `
UPDATE tmpusers
SET token = $1
WHERE token = $2
`

// ForwardToken will try to update a user based on an old token
// If the user doesn't exist in the temp database, or the token is
// invalid, BADTOKEN will be returned
func (d *Database) ForwardToken(ctx context.Context, old string) string {
	defer delete(d.valid, old)

	t, ok := d.valid[old]
	if !ok || t.Before(time.Now()) {
		return "BADTOKEN"
	}

	token := <-d.Temps

	//res, err := d.db.ExecContext(ctx, updateToken, token, old)
	//if err != nil { return "BADTOKEN" }
	//if row, e := res.RowsAffected(); e != nil || row < 1 { return "BADTOKEN" }
	if _, e := d.db.ExecContext(ctx, updateToken, token, old); e != nil {
		return "BADTOKEN"
	}

	return token
}

const findUserTemp = `
SELECT username, email, country, password
FROM tmpusers
WHERE token = $1
`

const deleteUserTemp = `
DELETE FROM tmpusers
WHERE token = $1
`

// Take in temp entry token
func (d *Database) UserFromTemp(ctx context.Context, token string) error {
	var email, password string
	user := User{
		Token: <-d.Tokens,
	}

	defer delete(d.valid, token)
	defer d.db.ExecContext(ctx, deleteUserTemp, token)

	t, ok := d.valid[token]
	if !ok || t.Before(time.Now()) {
		return errors.New("invalid/expired token")
	}

	rows, err := d.db.QueryContext(ctx, findUserTemp, token)
	if err != nil {
		return err
	}

	for rows.Next() {
		rows.Scan(&user.Username, &email, &user.Country, &password)

		// tokens are unique, so take the first and run away laughing
		return d.CreateUser(ctx, &user, email, password)
	}

	return errors.New("invalid/expired signup session")
}
