package ioutils

import "sync"

type Buffer[T any] struct {
	callNew func() T
	pool    sync.Pool
}

func NewBuffer[T any](callNew func() T) *Buffer[T] {
	return &Buffer[T]{
		pool: sync.Pool{New: func() any { return callNew() }},
	}
}

func (v *Buffer[T]) Get() T {
	buf, ok := v.pool.Get().(T)
	if !ok {
		buf = v.callNew()
	}
	return buf
}

func (v *Buffer[T]) Put(t T) {
	v.pool.Put(t)
}
