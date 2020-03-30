package template

import (
	"testing"
	"text/template"
)

func TestTemplates(t *testing.T) {
	if _, e := template.ParseGlob("*.tpl"); e != nil {
		t.Error(e)
	}

	if _, e := template.ParseGlob("help/*.tpl"); e != nil {
		t.Error(e)
	}

	if _, e := template.ParseGlob("doctor/*.tpl"); e != nil {
		t.Error(e)
	}

	if _, e := template.ParseGlob("patient/*.tpl"); e != nil {
		t.Error(e)
	}

	if _, e := template.ParseGlob("plugins/*.tpl"); e != nil {
		t.Error(e)
	}
}
