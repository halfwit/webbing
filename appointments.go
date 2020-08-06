package plugins

import (
	"github.com/olmaxmedical/router"
)

// AddAppointment registers an appointment into the appointment book
// TODO(halfwit) message/email client to fill out Symptoms form
const AddAppointment router.PluginMask = 1 << 1

func init() {
	b := &router.Plugin{
		Name:     "Add Appointments",
		Run:      nil,
		Validate: addAppt,
	}
	router.AddPlugin(b, AddAppointment)
}

func addAppt(s *router.Request) error {
	// get &id=
	// Auth user and requisite key
	// TODO(halfwit) create unique key in patient/appointments 
	return nil
}
