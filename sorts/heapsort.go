/*
 *  Copyright (c) 2019-2024 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

// see: https://en.wikipedia.org/wiki/Heapsort

package sorts

func Heapsort[T any](list []T, less func(i, j int) bool) {
	for i := len(list)/2 - 1; i >= 0; i-- {
		heapsortSort[T](list, i, len(list), less)
	}
	for i := len(list) - 1; i >= 0; i-- {
		list[0], list[i] = list[i], list[0]
		heapsortSort[T](list, 0, i, less)
	}
}

func heapsortSort[T any](list []T, parent, max int, less func(i, j int) bool) {
	for {
		child := parent
		left, right := parent*2+1, parent*2+2
		if left < max && !less(left, child) {
			child = left
		}
		if right < max && !less(right, child) {
			child = right
		}
		if child == parent {
			return
		}
		list[parent], list[child] = list[child], list[parent]
		parent = child
	}
}
