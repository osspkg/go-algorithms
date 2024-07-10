package ioutils

import "sync"

type TBuffer interface {
	Reset()
}

type Buffer[T TBuffer] struct {
	callNew func() T
	pool    sync.Pool
}

func NewBuffer[T TBuffer](callNew func() T) *Buffer[T] {
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
	t.Reset()
	v.pool.Put(t)
}

type SliceBuffer[T any] struct {
	B []T
}

func (v *SliceBuffer[T]) Reset() {
	v.B = v.B[:0]
}

func NewSliceBuffer[T any](l, c int) *Buffer[*SliceBuffer[T]] {
	return NewBuffer(func() *SliceBuffer[T] {
		return &SliceBuffer[T]{B: make([]T, l, c)}
	})
}
