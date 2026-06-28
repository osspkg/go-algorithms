/*
 *  Copyright (c) 2019-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package token

import (
	"errors"
	"math/rand/v2"
	"sync"
	"time"
)

var (
	poolRnd = sync.Pool{New: func() any {
		return createRand()
	}}
	poolDigest = sync.Pool{New: func() any {
		return createDigest(createRand(), 128)
	}}
)

type digest struct {
	D []byte
}

func createRand() *rand.Rand {
	seed2 := uint64(time.Now().UnixNano())
	return rand.New(rand.NewPCG(seed2/100, seed2)) //nolint:gosec
}

func createDigest(rnd *rand.Rand, n int) *digest {
	b := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		v := rnd.IntN(255)
		b = append(b, byte(v))
	}
	return &digest{D: b}
}

var (
	table        []byte
	reverseTable [255]int
)

func init() {
	if err := SetTable("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"); err != nil {
		panic(err)
	}
}

func SetTable(s string) error {
	if len(s)%2 != 0 {
		return errors.New("token: table must be a multiple of 2")
	}
	if len(s) < 16 {
		return errors.New("token: table must be 16 or more in length")
	}
	if len(s) > 255 {
		return errors.New("token: table should be no longer than 255")
	}

	uniq := make(map[byte]struct{}, len(s))

	t, r := []byte(s), [255]int{}
	for i := 0; i < len(r); i++ {
		r[i] = -1
	}
	for i, b := range t {
		if _, ok := uniq[b]; ok {
			return errors.New("token: duplicate chars in table")
		}
		uniq[b] = struct{}{}
		r[b] = i
	}
	table, reverseTable = t, r
	return nil
}
