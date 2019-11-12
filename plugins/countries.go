package plugins

import (
	"sort"
	"strings"

	"github.com/olmaxmedical/olmax_go/router"
	"github.com/pariz/gountries"
	"golang.org/x/text/message"
)

// ListCountries - Populate a localized spinner to select country
const ListCountries router.PluginMask = 4

// Country - Mapping token to internationalized country code
type Country struct {
	ID   string
	Name string
}

type countries struct {
	list []gountries.Country
}

var cache *countries

func init() {
	var l []gountries.Country
	c := gountries.New()
	for _, c := range c.FindAllCountries() {
		l = append(l, c)
	}
	cache = &countries{
		list: l,
	}
	sort.Sort(cache)
	b := &router.Plugin{
		Name:     "countrylist",
		Run:      Countries,
		Validate: nil,
	}
	router.AddPlugin(b, ListCountries)
}

// Len - For Sort implementation
func (c *countries) Len() int {
	return len(c.list)
}

// Less - For Sort implementation
func (c *countries) Less(i, j int) bool {
	switch strings.Compare(c.list[i].Name.Common, c.list[j].Name.Common) {
	case -1:
		return true
	default:
		return false
	}
}

// Swap - For Sort implementation
func (c *countries) Swap(i, j int) {
	tmp := c.list[i]
	c.list[i] = c.list[j]
	c.list[j] = tmp
}

// Countries - return a localized list of countries
func Countries(_ *router.Request) map[string]interface{} {
	// TODO(halfwit): Use Request to get a localized country name
	c := make(map[string]interface{})
	for _, item := range cache.list {
		c[item.Name.Common] = item.Name.Common
	}
	return c
}

// TODO: Export this so it's available to form parsing as a bitmask
func validateCountries(p *message.Printer, countries []string) string {
	for _, c := range countries {
		if msg := validateCountry(p, c); msg != "" {
			return msg
		}
	}
	return ""
}

func validateCountry(p *message.Printer, country string) string {
	for _, item := range cache.list {
		if item.Name.Common == country {
			return ""
		}
	}
	return p.Sprint("No country entered/nil value entered")
}
