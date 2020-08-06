package router

import (
	"fmt"
)

// PluginMask - (Must be unique) ID for a plugin
type PluginMask uint32

// DEAD is a magic string to indicate a non-unique plugin key
const DEAD PluginMask = 1

var pluginCache map[PluginMask]*Plugin
var pluginKey []PluginMask

// Plugin - Provide extra data or functionality from GET/POST pages
type Plugin struct {
	Name     string
	Run      func(p *Request) map[string]interface{}
	Validate func(p *Request) error
}

func init() {
	pluginCache = make(map[PluginMask]*Plugin)
}

// ValidatePlugins - Make sure that its mapping isn't redundant with any other
// Plugins have external testing to validate they are correct
func ValidatePlugins() []error {
	errs := []error{}
	for key, item := range pluginCache {
		if item.Validate == nil {
			continue
		}
		if (key & DEAD) != 0 {
			errs = append(errs, fmt.Errorf("Error registering %s: Key requested already in use", item.Name))
		}
	}
	return errs
}

// AddPlugin - Add Plugin to map by key
func AddPlugin(p *Plugin, key PluginMask) {
	if pluginCache[key] != nil {
		key |= DEAD
	}
	pluginKey = append(pluginKey, key)
	pluginCache[key] = p
}
