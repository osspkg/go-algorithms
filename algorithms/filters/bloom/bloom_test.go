/*
 *  Copyright (c) 2019-2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package bloom

import (
	"testing"

	"go.osspkg.com/x/test"
)

func TestUnit_Bloom(t *testing.T) {
	bf, err := New(1000, 0.00001)
	test.NoError(t, err)

	bf.Add([]byte("hello"))
	bf.Add([]byte("user"))
	bf.Add([]byte("home"))

	test.False(t, bf.Contain([]byte("users")))
	test.True(t, bf.Contain([]byte("user")))
}

func TestUnit_Bloom2(t *testing.T) {
	_, err := New(0, 0.00001)
	test.Error(t, err)

	_, err = New(1, 1)
	test.Error(t, err)

	_, err = New(1, 0.0001)
	test.NoError(t, err)
}

func Benchmark_Bloom(b *testing.B) {
	bf, err := New(1000, 0.00001)
	if err != nil {
		b.FailNow()
	}
	bf.Add([]byte("hello"))
	bf.Add([]byte("user"))
	bf.Add([]byte("home"))

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		bf.Contain([]byte("users"))
	}
}
