package router

import (
	"sort"
	"strings"

	"github.com/pariz/gountries"
)

type Country struct {
	ID   string
	Name string
}

type countries struct {
	list []gountries.Country
}

var country *countries

func init() {
	var l []gountries.Country
	c := gountries.New()
	for _, c := range c.FindAllCountries() {
		l = append(l, c)
	}
	country = &countries{
		list: l,
	}
	sort.Sort(country)
}

func (c *countries) Len() int {
	return len(c.list)
}

func (c *countries) Less(i, j int) bool {
	switch strings.Compare(c.list[i].Name.Common, c.list[j].Name.Common) {
	case -1:
		return true
	default:
		return false
	}
}

func (c *countries) Swap(i, j int) {
	tmp := c.list[i]
	c.list[i] = c.list[j]
	c.list[j] = tmp
}

// TODO: Filter out any countries we don't support
func listcountries() []Country {
	var c []Country
	for _, item := range country.list {
		c = append(c, Country{
			ID:   item.Name.Common,
			Name: item.Name.Common,
		})
	}
	return c
}
