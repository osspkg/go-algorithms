/*
 *  Copyright (c) 2019-2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

// see: https://en.wikipedia.org/wiki/Cocktail_shaker_sort

package sorts

func Cocktail[T any](list []T, less func(i, j int) bool) {
	changed, toRight := false, true
	min, max := 0, len(list)-1
	for {
		changed = false
		if toRight {
			for i := min; i < max; i++ {
				if less(i+1, i) {
					list[i], list[i+1] = list[i+1], list[i]
					changed = true
				}
			}
			max--
		} else {
			for i := max; i >= min; i-- {
				if less(i+1, i) {
					list[i], list[i+1] = list[i+1], list[i]
					changed = true
				}
			}
			min++
		}
		toRight = !toRight
		if !changed || min >= max {
			return
		}
	}
}
