// Copyright 2015 Rick Beton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package date

import (
	"database/sql/driver"
	"testing"
)

func TestDateScan(t *testing.T) {
	cases := []struct {
		v        interface{}
		expected PeriodOfDays
	}{
		{PeriodOfDays(10000).Date().Local(), 10000},
	}

	for i, c := range cases {
		r := new(Date)
		e := r.Scan(c.v)
		if e != nil {
		}
		if r.DaysSinceEpoch() != c.expected {
			t.Errorf("%d: Got %v, want %d", i, *r, c.expected)
		}

		var d driver.Valuer = *r

		_, e = d.Value()
		if e != nil {
			t.Errorf("%d: Got %v for %d", i, e, c.expected)
		}
	}
}

func TestDateScanWithJunk(t *testing.T) {
	cases := []struct {
		v        interface{}
		expected string
	}{
		{true, "bool true is not a meaningful date"},
		{true, "bool true is not a meaningful date"},
	}

	for i, c := range cases {
		r := new(Date)
		e := r.Scan(c.v)
		if e.Error() != c.expected {
			t.Errorf("%d: Got %q, want %q", i, e.Error(), c.expected)
		}
	}
}

func TestDateScanWithNil(t *testing.T) {
	var r *Date
	e := r.Scan(nil)
	if e != nil {
		t.Errorf("Got %v", e)
	}
	if r != nil {
		t.Errorf("Got %v", r)
	}
}
