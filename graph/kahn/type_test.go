/*
 *  Copyright (c) 2019-2023 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package kahn

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnit_KahnCoherent(t *testing.T) {
	graph := New()

	require.NoError(t, graph.Add("a", "b"))
	require.NoError(t, graph.Add("a", "c"))
	require.NoError(t, graph.Add("a", "d"))
	require.NoError(t, graph.Add("a", "e"))
	require.NoError(t, graph.Add("b", "d"))
	require.NoError(t, graph.Add("c", "d"))
	require.NoError(t, graph.Add("c", "e"))
	require.NoError(t, graph.Add("d", "e"))

	require.NoError(t, graph.Build())

	result := graph.Result()
	require.True(t, len(result) == 5)

	require.Contains(t, []string{"a,b,c,d,e", "a,c,b,d,e"}, strings.Join(result, ","))
}

func TestUnit_KahnCoherentBreakPoint(t *testing.T) {
	graph := New()

	require.NoError(t, graph.Add("a", "b"))
	require.NoError(t, graph.Add("a", "c"))
	require.NoError(t, graph.Add("a", "d"))
	require.NoError(t, graph.Add("a", "e"))
	require.NoError(t, graph.Add("b", "d"))
	require.NoError(t, graph.Add("c", "d"))
	require.NoError(t, graph.Add("c", "e"))
	require.NoError(t, graph.Add("d", "e"))

	graph.BreakPoint("c")

	require.NoError(t, graph.Build())

	result := graph.Result()
	require.True(t, len(result) <= 3)

	require.Contains(t, []string{"a,b,c", "a,c"}, strings.Join(result, ","))
}

func TestUnit_KahnCyclical(t *testing.T) {
	graph := New()

	require.NoError(t, graph.Add("1", "2"))
	require.NoError(t, graph.Add("2", "3"))
	require.NoError(t, graph.Add("3", "2"))

	require.Error(t, graph.Build())
}

func Benchmark_Kahn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			graph := New()
			_ = graph.Add("1", "2") //nolint:errcheck
			_ = graph.Add("1", "3") //nolint:errcheck
			_ = graph.Add("3", "4") //nolint:errcheck
			_ = graph.Add("2", "4") //nolint:errcheck
			_ = graph.Add("4", "5") //nolint:errcheck
			_ = graph.Build()       //nolint:errcheck
		}()
	}
}
