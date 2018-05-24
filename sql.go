// Copyright 2015 Rick Beton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package date

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// These methods allow Date and PeriodOfDays to be fields stored in an
// SQL database by implementing the database/sql/driver interfaces.
// The underlying column type is simply an integer.

// Scan parses some value. It implements sql.Scanner,
// https://golang.org/pkg/database/sql/#Scanner
func (d *Date) Scan(value interface{}) (err error) {
	if value == nil {
		return nil
	}

	switch value.(type) {
	case time.Time:
		*d = NewAt(value.(time.Time))
	default:
		err = fmt.Errorf("%T %+v is not a meaningful date", value, value)
	}
	return
}

// Value converts the value to an string. It implements driver.Valuer,
// https://golang.org/pkg/database/sql/driver/#Valuer
func (d Date) Value() (driver.Value, error) {
	return d.String(), nil
}
