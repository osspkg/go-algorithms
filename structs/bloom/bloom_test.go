/*
 *  Copyright (c) 2019-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package bloom_test

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash"
	"hash/fnv"
	"testing"

	"github.com/cespare/xxhash/v2"
	"go.osspkg.com/casecheck"

	"go.osspkg.com/algorithms/structs/bloom"
)

func TestUnit_Bloom(t *testing.T) {
	bf, err := bloom.New(bloom.Quantity(1000, 0.00001))
	casecheck.NoError(t, err)

	bf.Add([]byte("hello"))
	bf.Add([]byte("user"))
	bf.Add([]byte("home"))

	casecheck.False(t, bf.Contain([]byte("users")))
	casecheck.True(t, bf.Contain([]byte("user")))
}

func TestUnit_Bloom2(t *testing.T) {
	_, err := bloom.New(bloom.Quantity(0, 0.00001))
	casecheck.Error(t, err)

	_, err = bloom.New(bloom.Quantity(1, 1))
	casecheck.Error(t, err)

	_, err = bloom.New(bloom.Quantity(1, 0.0001))
	casecheck.NoError(t, err)
}

func runBloom(b *testing.B, size uint64, rate float64, h func() hash.Hash) {
	bf, err := bloom.New(bloom.Quantity(size, rate), bloom.HashFunc(h))
	if err != nil {
		b.FailNow()
	}

	b.ResetTimer()
	b.ReportAllocs()
	b.ReportMetric(float64(bf.SizeBytes()/(1024*1024)), "Mb")
	b.ReportMetric(float64(bf.MaxElements()), "elm")

	var i int
	for i = 0; i < b.N; i++ {
		bf.Add([]byte(fmt.Sprintf("u%d", i)))
		if !bf.Contain([]byte(fmt.Sprintf("u%d", i))) {
			b.Log(i)
			b.FailNow()
		}
	}

	if bf.Contain([]byte(fmt.Sprintf("u%d", i+1))) {
		b.Log(i + 1)
		b.FailNow()
	}
}

const (
	vSize = 10_000_000
	vRate = 0.0001
)

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
