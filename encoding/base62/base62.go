/*
 *  Copyright (c) 2019-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package base62

import "go.osspkg.com/algorithms/sorts"

const size = 62

type Base62 struct {
	enc []byte
	dec map[byte]uint64
}

func New(alphabet string) *Base62 {
	if len(alphabet) != size {
		panic("encoding alphabet is not 62-bytes long")
	}
	v := &Base62{
		enc: []byte(alphabet),
		dec: make(map[byte]uint64, size),
	}
	for i, b := range v.enc {
		v.dec[b] = uint64(i)
	}
	return v
}

func (v *Base62) Encode(id uint64) string {
	result := make([]byte, 0, 11)
	for id > 0 {
		result = append(result, v.enc[id%size])
		id /= size
	}
	sorts.Reverse(result)
	return string(result)
}

func (v *Base62) Decode(data string) uint64 {
	var id uint64
	for _, b := range []byte(data) {
		id = id*size + v.dec[b]
	}
	return id
}
