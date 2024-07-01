/*
 *  Copyright (c) 2019-2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package aesgcm_test

import (
	"testing"

	"go.osspkg.com/x/encryption/aesgcm"
	"go.osspkg.com/x/random"
	"go.osspkg.com/x/test"
)

func TestUnit_Codec(t *testing.T) {
	rndKey := random.String(32)
	message := []byte("Hello World!")

	c, err := aesgcm.New(rndKey)
	test.NoError(t, err)

	enc1, err := c.Encrypt(message)
	test.NoError(t, err)

	dec1, err := c.Decrypt(enc1)
	test.NoError(t, err)

	test.Equal(t, message, dec1)

	c, err = aesgcm.New(rndKey)
	test.NoError(t, err)

	enc2, err := c.Encrypt(message)
	test.NoError(t, err)

	test.NotEqual(t, enc1, enc2)

	dec2, err := c.Decrypt(enc1)
	test.NoError(t, err)

	test.Equal(t, message, dec2)

}
