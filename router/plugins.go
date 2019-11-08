package router

import (
	"fmt"

	"golang.org/x/text/message"
)

// DEAD is a magic string to indicate a non-unique plugin key
const DEAD = 666

var pluginCache map[int]*Plugin

// Plugin - Provide extra data or functionality from GET/POST pages
type Plugin struct {
	Name     string
	Run      func(p *message.Printer) map[string]interface{}
	Validate func() error
}

func init() {
	pluginCache = make(map[int]*Plugin)
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
		if (key & DEAD) != 0 {
			errs = append(errs, fmt.Errorf("Key requested already in use for plugin %s: %d", item.Name, key|DEAD))
		}
	}
	return errs
}

// AddPlugin - Add Plugin to map by key
func AddPlugin(p *Plugin, key int) {
	if pluginCache[key] != nil {
		key &= DEAD
	}
	pluginCache[key] = p
}
