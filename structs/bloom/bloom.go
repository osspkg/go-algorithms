/*
 *  Copyright (c) 2019-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

// see: https://en.wikipedia.org/wiki/Bloom_filter

// Фильтр Блума - это пространственно-эффективная вероятностная структура данных,
// созданная для проверки наличия элемента в множестве. Он спроектирован невероятно
// быстрым при минимальном использовании памяти ценой потенциальных ложных срабатываний.
// Существует возможность получить ложноположительное срабатывание
// (элемента в множестве нет, но структура данных сообщает, что он есть),
// но не ложноотрицательное. Другими словами, очередь возвращает или "возможно в наборе", или "определённо не в наборе".
// Фильтр Блума может использовать любой объём памяти, однако чем он больше, тем меньше вероятность ложного срабатывания.

package bloom

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"hash"
	"math"
	"sync"

	"github.com/cespare/xxhash/v2"
)

const blockSize = 64

type Bloom struct {
	bits  []uint64
	size  uint64
	salts [][64]byte

	optSize uint64
	optRate float64

	pool *sync.Pool
	mux  sync.RWMutex
}

type Option func(b *Bloom)

func HashFunc(h func() hash.Hash) Option {
	return func(b *Bloom) {
		b.pool = &sync.Pool{New: func() any { return h() }}
	}
}

func Quantity(size uint64, rate float64) Option {
	return func(b *Bloom) {
		b.optSize = size
		b.optRate = rate
	}
}

func New(opts ...Option) (*Bloom, error) {
	b := &Bloom{
		optSize: 10_000_000,
		optRate: 0.001,
		pool:    &sync.Pool{New: func() any { return xxhash.New() }},
	}

	for _, opt := range opts {
		opt(b)
	}

	if b.optSize == 0 {
		return nil, fmt.Errorf("bitset size cannot be 0")
	}
	if b.optRate <= 0 || b.optRate >= 1.0 {
		return nil, fmt.Errorf("false positive rate must be between 0 and 1")
	}

	m, k := b.calcOptimalParams(b.optSize, b.optRate)

	b.size = m
	b.bits = make([]uint64, m/blockSize+1)
	b.salts = make([][64]byte, k)

	for i := 0; i < int(k); i++ {
		if _, err := rand.Read(b.salts[i][:]); err != nil {
			return nil, fmt.Errorf("generate hash salt: %w", err)
		}
	}

	return b, nil
}

func (b *Bloom) MaxElements() uint64 {
	b.mux.RLock()
	defer b.mux.RUnlock()

	return b.size
}

func (b *Bloom) SizeBytes() int {
	b.mux.RLock()
	defer b.mux.RUnlock()

	return len(b.bits) * 8
}

func (b *Bloom) Add(v []byte) {
	b.mux.Lock()
	defer b.mux.Unlock()

	for i := 0; i < len(b.salts)-1; i++ {
		p := b.createHash(i, v)
		index, num := b.getIndex(p)
		b.bits[index] |= num
	}
}

func (b *Bloom) Contain(v []byte) bool {
	b.mux.RLock()
	defer b.mux.RUnlock()

	for i := 0; i < len(b.salts)-1; i++ {
		p := b.createHash(i, v)
		index, num := b.getIndex(p)
		if b.bits[index]&num > 0 {
			continue
		}
		return false
	}
	return true
}

func (b *Bloom) createHash(saltIndex int, key []byte) uint64 {
	h, ok := b.pool.Get().(hash.Hash)
	if !ok {
		panic("failed get hash function from pool")
	}
	defer func() {
		h.Reset()
		b.pool.Put(h)
	}()

	h.Write(key)
	h.Write(b.salts[saltIndex][:])

	return binary.BigEndian.Uint64(h.Sum(nil)) % b.size
}

func (*Bloom) getIndex(p uint64) (uint64, uint64) {
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
