package plugins

import "github.com/olmaxmedical/router"

// Search - generic search for doctors in area
const Search router.PluginMask = 1 << 13

func init() {
	b := &router.Plugin{
		Name:     "search",
		Run:      nil,
		Validate: search,
	}
	router.AddPlugin(b, Search)
}

// Stuuuuubbb
func search(r *router.Request) error {
	return nil
}
