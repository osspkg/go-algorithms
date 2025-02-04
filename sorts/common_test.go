/*
 *  Copyright (c) 2019-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package sorts

import (
	"sort"
	"testing"
)

func Benchmark_Default_SortSlice(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		arr := []int{45, 61, 87, 20, 65, 36, 25, 86, 64, 4, 36, 53, 17, 38, 48, 52, 53, 59, 80, 79, 95, 72, 85, 52, 9, 12, 9, 36, 47, 34}
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
	}
}
