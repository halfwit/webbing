package db

import (
	"errors"
	"log"
)
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
	data["foo"] = &entry {
		first: "foo",
		last:  "bar",
		email: "foo@bar.com",
		pass:  "1234567890",
		role:  PatientAuth,
	}
	data["bar"] = &entry {
		first: "foo",
		last: "bar",
		email: "doc@bar.com",
		pass: "1234567890",
		role: DoctorAuth,
	}
}

// Placeholder code for database lookups
func CreateTempEntry(first, last, email, pass, token string) {
	tmpdata[token] = &entry{
		first: first,
		last:  last,
		email: email,
		pass:  pass,
	}
}

func RemoveTempEntry(token string) {
	delete(tmpdata, token)
}

func FindTempEntry(token string) bool {
	if _, ok := tmpdata[token]; ok {
		return true
	}
	return false
}

func CreateEntry(token string) {
	log.Println(data)
	if ent, ok := tmpdata[token]; ok {
		data[token] = &entry {
			first: ent.first,
			last:  ent.last,
			email: ent.email,
			pass:  ent.pass,
			role:  PatientAuth,
		}
		delete(tmpdata, token)
	}
}

type User struct {
	First string
	Last  string
	Email string
}

// Dummy function of all dummy functions, this will eventually call cookie functions to maintain forward secrecy
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

func UpdateToken(old, new string) bool {
	defer delete(data, old)
	if ent, ok := data[old]; ok {
		data[new] = ent
		return true
	}
	return false
}

func FindEntry(token string) bool {
	if _, ok := data[token]; ok {
		return true
	}
	return false
}

func ValidateLogin(username, password string) bool {
	for _, client := range data {
		if client.email == username && client.pass == password {
			return true
		}
	}
	return false
}

func UserRole(username string) Access {
	for _, client := range data {
		if client.email != username {
			continue
		}
		return client.role
	}
	return GuestAuth
}

func UserExists(email string) bool {
	for _, client := range data {
		if client.email == email {
			return true
		}
	}
	return false
}

func UpdateUserPassword(token, pass string) {
	if _, ok := data[token]; ! ok {
		return
	}
	data[token].pass = pass
}
