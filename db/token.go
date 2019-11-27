package db

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/olmaxmedical/olmax_go/session"
)

// NewToken returns a unique token string
func NewToken() string {
	u, _ := uuid.NewRandom()
	t := u.String()
	return t
}

// ValidateToken - verify old token was correct, set new
func ValidateToken(h *http.Request, s session.Session) error {
	defer s.Delete("token")
	if h == nil {
		return errors.New("Invalid session")
	}
	token := h.PostFormValue("token")
	if s.Get("token") != token {
		return errors.New("Invalid/expired token")
	}
	s.Set("token", NewToken())
	return nil
}
