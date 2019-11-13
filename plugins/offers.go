package plugins

import "github.com/olmaxmedical/olmax_go/router"

// Offer - Request a time slot with doctor
const MakeOffer router.PluginMask = 8

func init() {
	b := &router.Plugin{
		Name:     "offer",
		Run:      nil,
		Validate: offer,
	}
	router.AddPlugin(b, MakeOffer)
}

func offer(s *router.Request) error {
	return nil
}
