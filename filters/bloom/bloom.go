/*
 *  Copyright (c) 2019-2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

// see: https://en.wikipedia.org/wiki/Bloom_filter

package bloom

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"math"
	"sync"
)

const blockSize = 64

type Bloom struct {
	bits  []uint64
	size  uint64
	salts [][32]byte

	mux sync.RWMutex
}

func New(n uint64, p float64) (*Bloom, error) {
	if n == 0 {
		return nil, fmt.Errorf("bitset size cannot be 0")
	}
	if p <= 0 || p >= 1.0 {
		return nil, fmt.Errorf("false positive rate must be between 0 and 1")
	}

	b := &Bloom{}
	m, k := b.calcOptimalParams(n, p)
	b.size = m
	b.bits = make([]uint64, m/blockSize+1)
	b.salts = make([][32]byte, k)
	for i := 0; i < int(k); i++ {
		if _, err := rand.Read(b.salts[i][:]); err != nil {
			return nil, fmt.Errorf("generate hash salt: %w", err)
		}
	}
	return b, nil
}

func (b *Bloom) Add(v []byte) {
	b.mux.Lock()
	defer b.mux.Unlock()

	for i := 0; i < len(b.salts)-1; i++ {
		p := b.createHash(i, v)
		indx, num := b.getPosition(p)
		b.bits[indx] |= num
	}
}

func (b *Bloom) Contain(v []byte) bool {
	b.mux.RLock()
	defer b.mux.RUnlock()

	for i := 0; i < len(b.salts)-1; i++ {
		p := b.createHash(i, v)
		indx, num := b.getPosition(p)
		if b.bits[indx]&num > 0 {
			continue
		}
		return false
	}
	return true
}

func (b *Bloom) createHash(i int, key []byte) uint64 {
	mac := hmac.New(sha1.New, b.salts[i][:])
	mac.Write(key)
	return binary.BigEndian.Uint64(mac.Sum(nil)) % b.size
}

func (*Bloom) getPosition(p uint64) (uint64, uint64) {
	index := uint64(math.Ceil(float64(p+1)/blockSize)) - 1
	num := uint64(1) << (p - index*blockSize)
	return index, num
}

func (*Bloom) calcOptimalParams(n uint64, p float64) (m, k uint64) {
	m = uint64(math.Ceil(-float64(n) * math.Log(p) / math.Pow(math.Log(2.0), 2.0)))
	if m == 0 {
		m = 1
	}
	k = uint64(math.Ceil(float64(m) * math.Log(2.0) / float64(n)))
	if k == 0 {
		k = 1
	}
	return
}
