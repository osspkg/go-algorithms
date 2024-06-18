/*
 *  Copyright (c) 2019-2024 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
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
		arr := []int{30, 29, 28, 27, 26, 25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
	}
}
