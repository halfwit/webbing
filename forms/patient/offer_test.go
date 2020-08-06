package forms

import (
	"net/url"
	"testing"

	"github.com/olmaxmedical/forms/util"
)

func TestOffer(t *testing.T) {
	values := url.Values{}

	values.Add("Amount", "0.1234")
	values.Add("startDate", "2020-04-04T00:00:00")
	values.Add("endDate", "2020-06-06T00:00:00")

	if e := util.TestValues(values, offer); e != nil {
		t.Error(e)
	}

	values.Set("Amount", "-1")
	if e := util.TestValues(values, offer); e == nil {
		t.Error("invalid BTC rate allowed")
	}

	values.Set("Amount", "0.1234")
	values.Set("startDate", "1995-30-30T23:23:59")
	if e := util.TestValues(values, offer); e == nil {
		t.Error("invalid date allowed")
	}
}
