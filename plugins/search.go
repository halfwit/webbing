package plugins

import "github.com/olmaxmedical/olmax_go/router"

// Search - generic search for doctors in area
const Search router.PluginMask = 11

func init() {
	b := &router.Plugin{
		Name:     "Search",
		Run:      nil,
		Validate: search,
	}
	router.AddPlugin(b, Search)
}

// Stuuuuubbb
func search(r *router.Request) error {
	return nil
}
