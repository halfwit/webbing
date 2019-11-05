package db

import (
	"errors"
	"log"
)

// Access - Who can access the data
type Access uint8

const (
	GuestAuth Access = 1 << iota
	PatientAuth
	DoctorAuth
)

type entry struct {
	first string
	last  string
	email string
	pass  string
	role  Access
}

// NOTE: stubs for database calls.
var tmpdata map[string]*entry
var data map[string]*entry

func init() {
	tmpdata = make(map[string]*entry)
	data = make(map[string]*entry)
	// NOTE: Dummy entry for testing. This goes away when we have a real db
	data["foo"] = &entry{
		first: "foo",
		last:  "bar",
		email: "foo@bar.com",
		pass:  "1234567890",
		role:  PatientAuth,
	}
	data["bar"] = &entry{
		first: "foo",
		last:  "bar",
		email: "doc@bar.com",
		pass:  "1234567890",
		role:  DoctorAuth,
	}
}

// CreateTempEntry - Temporary (time limited) database entry
func CreateTempEntry(first, last, email, pass, token string) {
	tmpdata[token] = &entry{
		first: first,
		last:  last,
		email: email,
		pass:  pass,
	}
}

// RemoveTempEntry - Called after timeout - this should be internal to db
func RemoveTempEntry(token string) {
	delete(tmpdata, token)
}

// FindTempEntry - validate entry still exists
func FindTempEntry(token string) bool {
	if _, ok := tmpdata[token]; ok {
		return true
	}
	return false
}

// CreateEntry - Add a permanent entry to the database
func CreateEntry(token string) {
	log.Println(data)
	if ent, ok := tmpdata[token]; ok {
		data[token] = &entry{
			first: ent.first,
			last:  ent.last,
			email: ent.email,
			pass:  ent.pass,
			role:  PatientAuth,
		}
		delete(tmpdata, token)
	}
}

// User - Any registered user on the site
type User struct {
	First string
	Last  string
	Email string
}

// FromCookie - look up by Cookie token
func FromCookie(token string) (*User, error) {
	if u, ok := data[token]; ok {
		return &User{
			First: u.first,
			Last:  u.last,
			Email: u.email,
		}, nil
	}
	return nil, errors.New("No such user")
}

// UpdateToken - Change entry status for temp entries
func UpdateToken(old, new string) bool {
	defer delete(data, old)
	if ent, ok := data[old]; ok {
		data[new] = ent
		return true
	}
	return false
}

// FindEntry - Look up if token is still valid
func FindEntry(token string) bool {
	if _, ok := data[token]; ok {
		return true
	}
	return false
}

// ValidateLogin - Dummy function for login
func ValidateLogin(username, password string) bool {
	for _, client := range data {
		if client.email == username && client.pass == password {
			return true
		}
	}
	return false
}

// UserRole - Find Access mappings for given user
func UserRole(username string) Access {
	for _, client := range data {
		if client.email != username {
			continue
		}
		return client.role
	}
	return GuestAuth
}

// UserExists - Look up by email if user is in db
func UserExists(email string) bool {
	for _, client := range data {
		if client.email == email {
			return true
		}
	}
	return false
}

// UpdateUserPassword - Dummy funtion to update password
func UpdateUserPassword(token, pass string) {
	if _, ok := data[token]; !ok {
		return
	}
	data[token].pass = pass
}
