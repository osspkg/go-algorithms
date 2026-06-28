/*
 *  Copyright (c) 2019-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package token

import (
	"errors"
)

func (t T64) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

func (t *T64) UnmarshalText(data []byte) error {
	nt, ok := ParseBytes(data)
	if !ok {
		return errors.New("invalid token")
	}
	*t = nt
	return nil
}

func (t T64) MarshalBinary() ([]byte, error) {
	return t.MarshalText()
}

func (t *T64) UnmarshalBinary(data []byte) error {
	return t.UnmarshalText(data)
}

func (t T64) MarshalJSON() ([]byte, error) {
	return t.MarshalText()
}

func (t *T64) UnmarshalJSON(data []byte) error {
	return t.UnmarshalText(data)
}
