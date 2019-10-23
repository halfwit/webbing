package router

import (
	"github.com/google/uuid"
)

var tokens []string

func newToken() string {
	u, _ := uuid.NewRandom()
	t := u.String()
	tokens = append(tokens, t)
	return t
}

func validateToken(token string) bool {
	for n, t := range tokens {
		if token == t {
			// n will always be at least 0, tokens at least 1
			tokens[n] = tokens[len(tokens)-1]
			tokens = tokens[:len(tokens)-1]
			return true
		}
	}
	return false
}
