/*
 *  Copyright (c) 2019-2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package sorts

func Reverse[T any](list []T) {
	j := len(list) - 1
	for i := 0; i < len(list)/2; i++ {
		list[i], list[j] = list[j], list[i]
		j--
	}
}
