package ioutils

import (
	"testing"

	"go.osspkg.com/x/test"
)

func TestUnit_Buffer(t *testing.T) {
	buf := NewSliceBuffer[byte](2, 10)

	item := buf.Get()
	test.True(t, len(item.B) == 2)
	test.True(t, cap(item.B) == 10)
	buf.Put(item)
}

func Benchmark_Buffer(b *testing.B) {
	buf := NewSliceBuffer[byte](0, 10)

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			item := buf.Get()
			if len(item.B) != 0 {
				b.FailNow()
			}
			item.B = append(item.B, []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}...)
			buf.Put(item)
		}
	})

}
