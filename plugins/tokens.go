package plugins

import (
	"github.com/google/uuid"
	"github.com/olmaxmedical/olmax_go/router"
)

var tokens []string

// SessionToken - An in-memory token to allow a client to track
const SessionToken router.PluginMask = 4

// FormToken - A database-persisted one time use token to relate forms to POST requests
const FormToken router.PluginMask = 5

func init() {
	b := &router.Plugin{
		Name:     "sessionToken",
		Run:      NewSessionToken,
		Validate: nil,
	}
	router.AddPlugin(b, SessionToken)
	c := &router.Plugin{
		Name:     "formToken",
		Run:      NewFormToken,
		Validate: nil,
	}
	router.AddPlugin(c, FormToken)
}

// NewSessionToken returns a unique session token
func NewSessionToken(r *router.Request) map[string]interface{} {
	return map[string]interface{}{
		"token": newToken(),
	}
}

// NewFormToken returns a unique token associated with a client's form entry session
// This will fall back to a database call
func NewFormToken(r *router.Request) map[string]interface{} {
	return map[string]interface{}{
		"token": newToken(),
	}
}

// TODO(halfwit) - Form plugins will use this
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

func newToken() string {
	u, _ := uuid.NewRandom()
	t := u.String()
	tokens = append(tokens, t)
	return t
}
