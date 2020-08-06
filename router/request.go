package router

import (
	"net/http"

	"github.com/olmaxmedical/database"
	"github.com/olmaxmedical/session"
	"golang.org/x/text/message"
)

// Request represents an incoming GET/POST
type Request struct {
	printer *message.Printer
	session session.Session
	request *http.Request
	user    string
	status  string
	path    string
	role    database.Access
}

// Printer - returns the client's localized printer handler
func (r *Request) Printer() *message.Printer {
	return r.printer
}

// Session - returns the client's session
func (r *Request) Session() session.Session {
	return r.session
}

// Request - underlying http.Request for forms and such
func (r *Request) Request() *http.Request {
	return r.request
}
