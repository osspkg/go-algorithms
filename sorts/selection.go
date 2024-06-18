/*
 *  Copyright (c) 2019-2024 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

// see: https://en.wikipedia.org/wiki/Selection_sort

package sorts

func Selection[T any](list []T, less func(i, j int) bool) {
	if len(list) < 2 {
		return
	}
	for i := 0; i < len(list)-1; i++ {
		min := i
		for j := i + 1; j < len(list); j++ {
			if less(j, min) {
				min = j
			}
		}
		list[i], list[min] = list[min], list[i]
	}
}
