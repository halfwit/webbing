package plugins

import (
	"github.com/olmaxmedical/olmax_go/router"
)

// AddAppointment registers an appointment into the appointment book
// TODO(halfwit) message/email client to fill out Symptoms form
const AddAppointment router.PluginMask = 1 << 15

func init() {
	b := &router.Plugin{
		Name:     "Add Appointments",
		Run:      nil,
		Validate: addAppt,
	}
	router.AddPlugin(b, AddAppointment)
}

func addAppt(s *router.Request) error {
	return nil
}
