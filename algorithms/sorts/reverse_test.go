/*
 *  Copyright (c) 2019-2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package sorts

import (
	"testing"

	"go.osspkg.com/x/test"
)

func TestUnit_Reverse(t *testing.T) {
	type testCase[T any] struct {
		name string
		list []T
		want []T
	}
	tests := []testCase[int]{
		{
			name: "case1",
			list: []int{2, 5, 3, 9, 0, 8},
			want: []int{8, 0, 9, 3, 5, 2},
		},
		{
			name: "case2",
			list: []int{2, 5, 3, 9, 0, 8, 4},
			want: []int{4, 8, 0, 9, 3, 5, 2},
		},
		{
			name: "case3",
			list: nil,
			want: nil,
		},
		{
			name: "case4",
			list: []int{2},
			want: []int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Reverse(tt.list)
			test.Equal(t, tt.want, tt.list)
		})
	}
}

func Benchmark_SortReverse(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		arr := []int{45, 61, 87, 20, 65, 36, 25, 86, 64, 4, 36, 53, 17, 38, 48, 52, 53, 59, 80, 79, 95, 72, 85, 52, 9, 12, 9, 36, 47, 34}
		Reverse(arr)
	}
}
