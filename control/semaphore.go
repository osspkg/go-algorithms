/*
 *  Copyright (c) 2019-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package control

type (
	Semaphore interface {
		Acquire()
		Release()
	}
	_semaphore struct {
		c chan struct{}
	}
)

func NewSemaphore(count uint64) Semaphore {
	return &_semaphore{
		c: make(chan struct{}, count),
	}
}

func (s *_semaphore) Acquire() {
	s.c <- struct{}{}
}

func (s *_semaphore) Release() {
	<-s.c
}
