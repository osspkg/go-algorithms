/*
 *  Copyright (c) 2019-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package token

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

func (t *T64) Scan(src interface{}) error {
	switch src := src.(type) {
	case nil:
		return nil

	case string:
		if src == "" {
			return nil
		}
		nt, ok := Parse(src)
		if !ok {
			return errors.New("scan: invalid token")
		}
		*t = nt

	case []byte:
		if len(src) == 0 {
			return nil
		}
		nt, ok := ParseBytes(src)
		if !ok {
			return errors.New("scan: invalid token")
		}
		*t = nt

	default:
		return fmt.Errorf("scan: unable to scan type %T into UUID", src)
	}

	return nil
}

func (t T64) Value() (driver.Value, error) {
	return t.String(), nil
}
