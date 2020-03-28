package plugins

import (
	"fmt"
	"sort"
	"strings"

	"github.com/olmaxmedical/router"
	"github.com/pariz/gountries"
	"golang.org/x/text/message"
)

// Countries - Populate a localized spinner to select country
const Countries router.PluginMask = 1 << 3

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
		Name:     "country",
		Run:      listCountries,
		Validate: validateCountries,
	}
	router.AddPlugin(b, Countries)
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

func listCountries(r *router.Request) map[string]interface{} {
	p := r.Printer()
	// TODO(halfwit): Use Request to get a localized country name
	c := make(map[string]interface{})
	c["label"] = p.Sprint("Select country/countries")
	for _, item := range cache.list {
		c[item.Name.Common] = item.Name.Common
	}
	return c
}

func validateCountries(r *router.Request) error {
	s := r.Request()
	for _, c := range s.PostFormValue("country") {
		if msg := checkCountry(r.Printer(), c); msg != nil {
			return msg
		}
	}
	return nil
}

func checkCountry(p *message.Printer, country rune) error {
	for _, item := range cache.list {
		if item.Name.Common == string(country) {
			return nil
		}
	}
	return fmt.Errorf("%s", p.Sprint("No country entered/nil value entered"))
}
