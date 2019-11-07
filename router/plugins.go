package router

import (
	"golang.org/x/text/message"
)

var pluginCache map[int]*Plugin

// Plugin - Provide extra data or functionality from GET/POST pages
type Plugin struct {
	Run func(p *message.Printer) map[string]interface{}
}

func init() {
	pluginCache = make(map[int]*Plugin)
}

// ValidatePlugins - Run through each plugin
// make sure that its mapping isn't redundant with any other
// Make sure the code runs accurately without error
func ValidatePlugins() []error {
	errs := []error{}
	return errs
}

// AddPlugin - Add Plugin to map by key
func AddPlugin(p *Plugin, key int) {
	pluginCache[key] = p
}
