/*
 *  Copyright (c) 2019-2024 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package shorten

type Shorten struct {
	toStr map[uint64]string
	toInt map[string]uint64
	len   uint64
}

func New(alphabet string) *Shorten {
	v := &Shorten{
		toStr: make(map[uint64]string),
		toInt: make(map[string]uint64),
		len:   uint64(len(alphabet)),
	}

	var i uint64
	for i = 0; i < v.len; i++ {
		v.toInt[alphabet[i:i+1]] = i
		v.toStr[i] = alphabet[i : i+1]
	}
	return v
}

func (v *Shorten) Encode(id uint64) string {
	s := ""
	for id > 0 {
		s = v.toStr[id%v.len] + s
		id /= v.len
	}
	return s
}

func (v *Shorten) Decode(data string) uint64 {
	var id, i uint64
	for i = 0; i < uint64(len(data)); i++ {
		id = id*v.len + v.toInt[data[i:i+1]]
	}
	return id
}
