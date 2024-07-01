/*
 *  Copyright (c) 2019-2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package random

import (
	"bytes"
	"fmt"
	"testing"
)

func TestUnit_Bytes(t *testing.T) {
	size := 10
	r1 := Bytes(size)
	r2 := Bytes(size)

	fmt.Println(string(r1), string(r2))

	if len(r1) != size || len(r2) != size {
		t.Errorf("invalid len, is not %d", size)
	}
	if bytes.Equal(r1, r2) {
		t.Errorf("result is not random")
	}
}

func TestUnit_BytesOf(t *testing.T) {
	size := 10
	src := []byte("1234567890")
	r1 := BytesOf(size, src)
	r2 := BytesOf(size, src)

	fmt.Println(string(r1), string(r2))

	if len(r1) != size || len(r2) != size {
		t.Errorf("invalid len, is not %d", size)
	}
	if bytes.Equal(r1, r2) {
		t.Errorf("result is not random")
	}
}

func Benchmark_Bytes64(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Bytes(64)
	}
}

func Benchmark_Bytes256(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Bytes(256)
	}
}
