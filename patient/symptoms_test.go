package forms

import (
	"net/url"
	"testing"
	"time"

	"github.com/olmaxmedical/forms/util"
)

func TestSymptoms(t *testing.T) {
	values := url.Values{}

	values.Add("bday", "1990-01-01T01:01:01")
	values.Add("onset", "2001-01-01T01:01:01")
	values.Add("gender", "male")
	values.Add("duration", "1")
	values.Add("reason", "test")
	values.Add("location", "test")
	values.Add("characteristic", "test")
	values.Add("aggreAlevi", "test")
	for _, i := range []string{
		"feversChills",
		"wtGainLoss",
		"vision",
		"lung",
		"heart",
		"bowel",
		"renal",
		"musSkel",
		"neuro",
		"psych",
	} {
		values.Add(i, "yes")
	}

	if e := util.TestValues(values, symptoms); e != nil {
		t.Error(e)
	}

	values.Set("bday", "1891-01-01T01:01:01")

	if e := util.TestValues(values, symptoms); e == nil {
		t.Error("forms parsing: invalid date accepted")
	}

	values.Set("bday", "1990-01-01T01:01:01")
	values.Set("onset", time.Now().Add(time.Hour+48).String())

	if e := util.TestValues(values, symptoms); e == nil {
		t.Error("form parsing: invalid onset accepted")
	}
}
