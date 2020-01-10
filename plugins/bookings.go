package plugins

import (
	"github.com/olmaxmedical/olmax_go/router"
)

// ListBookings retreives a list of all upcoming bookings for a given doctor
const ListBookings router.PluginMask = 1 << 2

func init() {
	b := &router.Plugin{
		Name:     "List Bookings",
		Run:      nil,
		Validate: listBookings,
	}
	router.AddPlugin(b, ListBookings)
}

func listBookings(s *router.Request) error {
	
	return nil
}
