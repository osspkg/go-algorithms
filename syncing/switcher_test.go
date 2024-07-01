/*
 *  Copyright (c) 2019-2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package syncing

import (
	"testing"

	"go.osspkg.com/x/test"
)

func TestNewSwitch(t *testing.T) {
	sync := NewSwitch()

	test.False(t, sync.IsOn())
	test.True(t, sync.IsOff())

	test.True(t, sync.On())
	test.False(t, sync.On())

	test.False(t, sync.IsOff())
	test.True(t, sync.IsOn())

}
