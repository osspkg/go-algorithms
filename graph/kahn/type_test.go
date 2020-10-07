package kahn

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_KahnCoherent(t *testing.T) {
	graph := New()

	assert.NoError(t, graph.Add("a", "b"))
	assert.NoError(t, graph.Add("a", "c"))
	assert.NoError(t, graph.Add("a", "d"))
	assert.NoError(t, graph.Add("a", "e"))
	assert.NoError(t, graph.Add("b", "d"))
	assert.NoError(t, graph.Add("c", "d"))
	assert.NoError(t, graph.Add("c", "e"))
	assert.NoError(t, graph.Add("d", "e"))

	assert.NoError(t, graph.Build())

	result := graph.Result()
	assert.True(t, len(result) > 0)

	assert.Contains(t, []string{"a,b,c,d,e", "a,c,b,d,e"}, strings.Join(result, ","))
}

func TestUnit_KahnCyclical(t *testing.T) {
	graph := New()

	assert.NoError(t, graph.Add("1", "2"))
	assert.NoError(t, graph.Add("2", "3"))
	assert.NoError(t, graph.Add("3", "2"))

	assert.Error(t, graph.Build())
}

func Benchmark_Kahn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			graph := New()
			_ = graph.Add("1", "2")
			_ = graph.Add("1", "3")
			_ = graph.Add("3", "4")
			_ = graph.Add("2", "4")
			_ = graph.Add("4", "5")
			_ = graph.Build()
		}()
	}
}
