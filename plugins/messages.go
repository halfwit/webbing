package plugins

import (
	"github.com/olmaxmedical/olmax_go/router"
)

// Messages - list all messages by key for use in a message template
const Messages router.PluginMask = 1 << 9

func init() {
	b := &router.Plugin{
		Name:     "messages",
		Run:      GetMessages,
		Validate: nil,
	}
	router.AddPlugin(b, Messages)
}

// GetMessages - return a map of message structs
func GetMessages(r *router.Request) map[string]interface{} {
	s := r.Session()
	if s == nil {
		return nil
	}
	return map[string]interface{}{
		// No, this won't actually do anything.
		"messages": s.Get("messages"),
	}
}
