/*
 *  Copyright (c) 2019-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package bitmap

import (
	"math"
	"testing"

	"go.osspkg.com/casecheck"
)

func TestUnit_Bitmap_calcBlockIndex(t *testing.T) {
	bm := New(65)

	for i := 0; i <= 65; i++ {
		bm.Set(uint64(i))

		casecheck.True(t, bm.Has(uint64(i)), "(1) for index: %d", i)
		casecheck.False(t, bm.Has(uint64(i+1)), "(2) for index: %d", i+1)
	}

	backup := bm.Dump()

	bm.Restore(make([]byte, len(backup)))

	for i := 0; i <= 65; i++ {
		casecheck.False(t, bm.Has(uint64(i)), "(3) for index: %d", i)
		casecheck.False(t, bm.Has(uint64(i+1)), "(4) for index: %d", i+1)
	}

	bm.Restore(backup)

	for i := 65; i >= 0; i-- {
		casecheck.True(t, bm.Has(uint64(i)), "(1) for index: %d", i)
		casecheck.False(t, bm.Has(uint64(i+1)), "(2) for index: %d", i+1)

		bm.Del(uint64(i))
		casecheck.False(t, bm.Has(uint64(i)), "(3) for index: %d", i+1)
	}

}

/*
goos: linux
goarch: amd64
pkg: go.osspkg.com/algorithms/structs/bitmap
cpu: 12th Gen Intel(R) Core(TM) i9-12900KF
Benchmark_Bitmap
Benchmark_Bitmap-4   	 9942369	       162.7 ns/op	       0 B/op	       0 allocs/op
*/
func Benchmark_Bitmap(b *testing.B) {
	index := uint64(math.MaxInt16)
	bm := New(index)

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			bm.Set(index)
			bm.Has(index)
			bm.Del(index)
		}
	})
}
