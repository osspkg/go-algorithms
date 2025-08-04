/*
 *  Copyright (c) 2019-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package bitmap

import (
	"sync"
)

const (
	blockSize = 8
)

type Bitmap struct {
	data    []byte
	size    uint64
	max     uint64
	mux     sync.RWMutex
	lockoff bool
}

type Option func(*Bitmap)

func DisableLock() Option {
	return func(o *Bitmap) {
		o.lockoff = true
	}
}

func New(maxIndex uint64, opts ...Option) *Bitmap {
	size := maxIndex/blockSize + maxIndex%blockSize

	bm := &Bitmap{
		max:  size * blockSize,
		size: size,
		data: make([]byte, size),
	}

	for _, opt := range opts {
		opt(bm)
	}

	return bm
}

func (b *Bitmap) CopyTo(dst *Bitmap) {
	if !b.lockoff {
		b.mux.Lock()
		defer b.mux.Unlock()
	}
	if !dst.lockoff {
		dst.mux.Lock()
		defer dst.mux.Unlock()
	}

	dst.data = make([]byte, len(b.data))
	copy(dst.data, b.data)
	dst.size = b.size
	dst.max = b.max
	dst.lockoff = b.lockoff
}

func (b *Bitmap) Set(index uint64) {
	if index > b.max {
		return
	}

	if !b.lockoff {
		b.mux.Lock()
		defer b.mux.Unlock()
	}

	b.data[index%b.size] |= 1 << (index % blockSize)
}

func (b *Bitmap) Del(index uint64) {
	if index > b.max {
		return
	}

	if !b.lockoff {
		b.mux.Lock()
		defer b.mux.Unlock()
	}

	b.data[index%b.size] &^= 1 << (index % blockSize)
}

func (b *Bitmap) Has(index uint64) bool {
	if index > b.max {
		return false
	}

	if !b.lockoff {
		b.mux.RLock()
		defer b.mux.RUnlock()
	}

	return (b.data[index%b.size] & (1 << (index % blockSize))) > 0
}

func (b *Bitmap) Dump() []byte {
	if !b.lockoff {
		b.mux.RLock()
		defer b.mux.RUnlock()
	}

	out := make([]byte, b.size)
	copy(out, b.data)
	return out
}

func (b *Bitmap) Restore(in []byte) {
	if !b.lockoff {
		b.mux.Lock()
		defer b.mux.Unlock()
	}

	b.data = make([]byte, len(in))
	copy(b.data, in)

	b.size = uint64(len(in))
	b.max = b.size * blockSize
}
