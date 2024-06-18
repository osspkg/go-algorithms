/*
 *  Copyright (c) 2019-2024 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

// see: https://en.wikipedia.org/wiki/Bubble_sort

package sorts

func Bubble[T any](list []T, less func(i, j int) bool) {
	var (
		changed bool
	)
	for j := 0; j < len(list)-1; j++ {
		changed = false
		for i := 0; i < len(list)-1-j; i++ {
			if less(i+1, i) {
				list[i], list[i+1] = list[i+1], list[i]
				changed = true
			}
		}
		if !changed {
			return
		}
	}
}
