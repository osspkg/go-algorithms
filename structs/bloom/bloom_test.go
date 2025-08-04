/*
 *  Copyright (c) 2019-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package bloom

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash"
	"hash/fnv"
	"reflect"
	"testing"

	"github.com/cespare/xxhash/v2"
	"go.osspkg.com/casecheck"
)

func TestUnit_Bloom(t *testing.T) {
	bf, err := New(Quantity(4, 0.01))
	casecheck.NoError(t, err)

	bf.Add("hello")
	bf.Add("user")
	bf.Add("home")

	casecheck.False(t, bf.Contain("users"))
	casecheck.True(t, bf.Contain("user"))
	casecheck.True(t, bf.Contain("hello"))
	casecheck.True(t, bf.Contain("home"))

	buf := bytes.NewBuffer(nil)
	casecheck.NoError(t, bf.Dump(buf))
	b1 := buf.Bytes()

	fmt.Println(string(b1))

	casecheck.NoError(t, bf.Restore(buf))
	buf = bytes.NewBuffer(nil)
	casecheck.NoError(t, bf.Dump(buf))
	b2 := buf.Bytes()

	casecheck.Equal(t, b1, b2)

	casecheck.False(t, bf.Contain("users"))
	casecheck.True(t, bf.Contain("user"))
	casecheck.True(t, bf.Contain("hello"))
	casecheck.True(t, bf.Contain("home"))
}

func TestUnit_Bloom2(t *testing.T) {
	_, err := New(Quantity(0, 0.00001))
	casecheck.Error(t, err)

	_, err = New(Quantity(1, 1))
	casecheck.Error(t, err)

	_, err = New(Quantity(1, 0.0001))
	casecheck.NoError(t, err)
}

func runBloom(b *testing.B, size uint64, rate float64, h func() hash.Hash) {
	bf, err := New(Quantity(size, rate), HashFunc(h))
	if err != nil {
		b.FailNow()
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		bf.Add(i)
		if !bf.Contain(i) {
			b.Fatal(i)
		}
	}
}

const (
	vSize = 10_000_000
	vRate = 0.1
)

/*
goos: linux
goarch: amd64
pkg: go.osspkg.com/algorithms/structs/bloom
cpu: 12th Gen Intel(R) Core(TM) i9-12900KF
Benchmark_Bloom_fnv128-4    	 2435511	       499.1 ns/op	     224 B/op	      19 allocs/op
Benchmark_Bloom_fnv128a-4   	 2512622	       486.0 ns/op	     224 B/op	      19 allocs/op
Benchmark_Bloom_md5-4       	 1253158	       942.2 ns/op	     160 B/op	      11 allocs/op
Benchmark_Bloom_sha1-4      	 1000000	      1077 ns/op	     224 B/op	      11 allocs/op
Benchmark_Bloom_sha256-4    	 1533015	       798.8 ns/op	     288 B/op	      11 allocs/op
Benchmark_Bloom_xxhash-4    	 3036788	       378.8 ns/op	      96 B/op	      11 allocs/op
*/

func Benchmark_Bloom_fnv128(b *testing.B) {
	runBloom(b, vSize, vRate, fnv.New128)
}

func Benchmark_Bloom_fnv128a(b *testing.B) {
	runBloom(b, vSize, vRate, fnv.New128a)
}

func Benchmark_Bloom_md5(b *testing.B) {
	runBloom(b, vSize, vRate, md5.New)
}

func Benchmark_Bloom_sha1(b *testing.B) {
	runBloom(b, vSize, vRate, sha1.New)
}

func Benchmark_Bloom_sha256(b *testing.B) {
	runBloom(b, vSize, vRate, sha256.New)
}

func Benchmark_Bloom_xxhash(b *testing.B) {
	runBloom(b, vSize, vRate, func() hash.Hash {
		return xxhash.New()
	})
}

func TestUnit_anyToBytes(t *testing.T) {
	tests := []struct {
		name string
		arg  any
		want []byte
	}{
		{
			name: "case Bytes",
			arg:  []byte("hello"),
			want: []byte("hello"),
		},
		{
			name: "case String",
			arg:  "hello",
			want: []byte("hello"),
		},
		{
			name: "case Int",
			arg:  12345,
			want: []byte{242, 192, 1},
		},
		{
			name: "case Struct",
			arg:  struct{ A int }{A: 1},
			want: []byte("{A:1}"),
		},
		{
			name: "case Ptr",
			arg:  &struct{ A int }{A: 1},
			want: []byte("&{A:1}"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := anyToBytes(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("anyToBytes() = %v, want %v", got, string(tt.want))
			}
		})
	}
}
