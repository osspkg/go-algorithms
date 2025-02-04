/*
 *  Copyright (c) 2019-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package base62

import (
	"fmt"
	"math"
	"testing"

	"go.osspkg.com/casecheck"
)

func TestEncode_EncodeDecode(t *testing.T) {
	tests := []struct {
		name string
		id   uint64
		want string
	}{
		{name: "Case1", id: 1, want: "p"},
		{name: "Case1", id: 2, want: "L"},
		{name: "Case1", id: 3, want: "K"},
		{name: "Case1", id: 4, want: "G"},
		{name: "Case1", id: 5, want: "R"},
		{name: "Case1", id: 6, want: "S"},
		{name: "Case1", id: 7, want: "u"},
		{name: "Case1", id: 8, want: "D"},
		{name: "Case1", id: 9, want: "v"},
		{name: "Case2", id: 10, want: "o"},
		{name: "Case3", id: 100, want: "pH"},
		{name: "Case4", id: 1000, want: "PD"},
		{name: "Case5", id: 10000, want: "LIn"},
		{name: "Case6", id: 100000, want: "c0k"},
		{name: "Case7", id: 1000000000, want: "pRmUWP"},
		{name: "Case8", id: 999999, want: "Glvp"},
		{name: "Case20", id: math.MaxUint64, want: "XNWjtpSoji4"},
	}

	v := New("0pLKGRSuDvorlO14Pjnd7XgQw9c8YhaIJ5iqtIHy3mWxM6C2TeAbFVBUkZfsNz")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := v.Encode(tt.id)
			casecheck.Equal(t, tt.want, h)
			id := v.Decode(h)
			casecheck.Equal(t, tt.id, id)
		})
	}
}

func TestEncode_Encode(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want uint64
	}{
		{name: "Case1", str: "a", want: 30},
		{name: "Case2", str: "aa", want: 1890},
		{name: "Case3", str: "aaaaaaaa", want: 107380379795850},
	}

	v := New("0pLKGRSuDvorlO14Pjnd7XgQw9c8YhaIJ5iqtIHy3mWxM6C2TeAbFVBUkZfsNz")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := v.Decode(tt.str)
			fmt.Println(h)
			casecheck.Equal(t, tt.want, h)
		})
	}
}

func Benchmark_base62(b *testing.B) {
	v := New("0pLKGRSuDvorlO14Pjnd7XgQw9c8YhaIJ5iqtIHy3mWxM6C2TeAbFVBUkZfsNz")

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			v.Decode(v.Encode(math.MaxUint64))
		}
	})
}
