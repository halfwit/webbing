package plugins

import (
	"errors"

	"github.com/google/uuid"
	"github.com/olmaxmedical/olmax_go/router"
)

//TODO(halfwit) Set up in memory tokens in db as well as
var tokens []string

// SessionToken - An in-memory token to allow a client to track
const SessionToken router.PluginMask = 1 << 13

// FormToken - A database-persisted one time use token to relate forms to POST requests
const FormToken router.PluginMask = 1 << 14

func init() {
	b := &router.Plugin{
		Name:     "sessionToken",
		Run:      NewSessionToken,
		Validate: ValidateToken,
	}
	router.AddPlugin(b, SessionToken)
	c := &router.Plugin{
		Name:     "formToken",
		Run:      NewFormToken,
		Validate: ValidateToken,
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
// TODO(halfwit) - database
func NewFormToken(r *router.Request) map[string]interface{} {
	return map[string]interface{}{
		"token": newToken(),
	}
}

// ValidateToken - Verify token exists
func ValidateToken(r *router.Request) error {
	s := r.Request()
	if s == nil {
		return errors.New("Invalid session")
	}
	token := s.PostFormValue("token")
	for n, t := range tokens {
		if token == t {
			// n will always be at least 0, tokens at least 1
			tokens[n] = tokens[len(tokens)-1]
			tokens = tokens[:len(tokens)-1]
			return nil
		}
	}
	return errors.New("Invalid/missing token")
}

func newToken() string {
	u, _ := uuid.NewRandom()
	t := u.String()
	tokens = append(tokens, t)
	return t
}
