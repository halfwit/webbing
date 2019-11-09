package router

import (
	"fmt"

	"golang.org/x/text/message"
)

// DEAD is a magic string to indicate a non-unique plugin key
const DEAD IncludeExtra = 0x0666000

var pluginCache map[IncludeExtra]*Plugin

// Plugin - Provide extra data or functionality from GET/POST pages
type Plugin struct {
	Name     string
	Run      func(p *message.Printer) map[string]interface{}
	Validate func() error
}

func init() {
	pluginCache = make(map[IncludeExtra]*Plugin)
}

// ValidatePlugins - Run through each plugin
// make sure that its mapping isn't redundant with any other
// Make sure the code runs accurately without error
func ValidatePlugins() []error {
	errs := []error{}
	for key, item := range pluginCache {
		err := item.Validate()
		if err != nil {
			errs = append(errs, err)
		}
		if (DEAD & key) != 0 {
			errs = append(errs, fmt.Errorf("Error registering %s: Key requested already in use (%d)", item.Name, key^DEAD))
		}
	}
	return errs
}

// AddPlugin - Add Plugin to map by key
func AddPlugin(p *Plugin, key IncludeExtra) {
	if pluginCache[key] != nil {
		key |= DEAD
	}
	pluginCache[key] = p
}
