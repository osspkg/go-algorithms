package kahn

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_KahnCoherent(t *testing.T) {
	kahn := New()

	assert.NoError(t, kahn.Add("a", "b"))
	assert.NoError(t, kahn.Add("a", "c"))
	assert.NoError(t, kahn.Add("a", "d"))
	assert.NoError(t, kahn.Add("a", "e"))
	assert.NoError(t, kahn.Add("b", "d"))
	assert.NoError(t, kahn.Add("c", "d"))
	assert.NoError(t, kahn.Add("c", "e"))
	assert.NoError(t, kahn.Add("d", "e"))

	assert.NoError(t, kahn.Build())

	result := kahn.Result()
	assert.True(t, len(result) > 0)

	assert.Contains(t, []string{"a,b,c,d,e", "a,c,b,d,e"}, strings.Join(result, ","))
}

func TestUnit_KahnCyclical(t *testing.T) {
	kahn := New()

	assert.NoError(t, kahn.Add("1", "2"))
	assert.NoError(t, kahn.Add("2", "3"))
	assert.NoError(t, kahn.Add("3", "2"))

	assert.Error(t, kahn.Build())
}

func Benchmark_Kahn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			kahn := New()
			_ = kahn.Add("1", "2")
			_ = kahn.Add("1", "3")
			_ = kahn.Add("3", "4")
			_ = kahn.Add("2", "4")
			_ = kahn.Add("4", "5")
			_ = kahn.Build()
		}()
	}
}
