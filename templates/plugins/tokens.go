package plugins

import (
	"github.com/olmaxmedical/database"
	"github.com/olmaxmedical/router"
)

// FormToken - A database-persisted one time use token to relate forms to POST requests
const FormToken router.PluginMask = 1 << 15

func init() {
	c := &router.Plugin{
		Name:     "formToken",
		Run:      newFormToken,
		Validate: validateToken,
	}
	router.AddPlugin(c, FormToken)
}

func validateToken(r *router.Request) error {
	return database.ValidateToken(r.Request(), r.Session())
}

// TODO(halfwit) - database
func newFormToken(r *router.Request) map[string]interface{} {
	return map[string]interface{}{
		"token": database.NewToken(),
	}
}
