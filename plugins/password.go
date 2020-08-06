package plugins

import (
	"errors"

	"github.com/olmaxmedical/database"
	"github.com/olmaxmedical/router"
)

// ValidateLogin - Check user/pass combo exists
const ValidateLogin router.PluginMask = 1 << 11

// ResetPassword - Update database entry for password
const ResetPassword router.PluginMask = 1 << 12

func init() {
	b := &router.Plugin{
		Name:     "login",
		Run:      nil,
		Validate: login,
	}
	router.AddPlugin(b, ValidateLogin)
	c := &router.Plugin{
		Name:     "setPassword",
		Run:      nil,
		Validate: setPass,
	}
	router.AddPlugin(c, ResetPassword)
}

func login(s *router.Request) error {
	r := s.Request()
	us := s.Session()
	p := s.Printer()
	user := r.PostFormValue("email")
	pass := r.PostFormValue("pass")
	if database.ValidateLogin(user, pass) {
		us.Set("username", user)
		us.Set("login", "true")
		us.Set("role", database.UserRole(user))
		us.Set("token", database.NewToken())
		return nil
	}
	return errors.New(p.Sprint("Invalid login"))
}

func setPass(s *router.Request) error {
	r := s.Request()
	p := s.Printer()

	pass := r.PostFormValue("password")
	repeat := r.PostFormValue("reenter")
	if pass != repeat {
		return errors.New(p.Sprint("Passwords do not match"))
	}
	token := r.PostFormValue("token")
	if !database.FindTempEntry(token) {
		return errors.New(p.Sprint("Session expired"))
	}
	database.UpdateUserPassword(token, pass)
	return nil
}
