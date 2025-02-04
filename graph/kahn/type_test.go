/*
 *  Copyright (c) 2019-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package kahn

import (
	"strings"
	"testing"

	"go.osspkg.com/casecheck"
)

func TestUnit_KahnCoherent(t *testing.T) {
	graph := New()
	graph.Add("a", "b")
	graph.Add("a", "c")
	graph.Add("a", "d")
	graph.Add("a", "e")
	graph.Add("b", "d")
	graph.Add("c", "d")
	graph.Add("c", "e")
	graph.Add("d", "e")
	casecheck.NoError(t, graph.Build())
	result := graph.Result()
	casecheck.True(t, len(result) == 5)
	casecheck.Equal(t, "a,b,c,d,e", strings.Join(result, ","))
}

func TestUnit_KahnCoherentBreakPoint(t *testing.T) {
	graph := New()
	graph.Add("a", "b")
	graph.Add("a", "c")
	graph.Add("a", "d")
	graph.Add("a", "e")
	graph.Add("b", "d")
	graph.Add("c", "d")
	graph.Add("c", "e")
	graph.Add("d", "e")
	graph.BreakPoint("d")
	casecheck.NoError(t, graph.Build())
	result := graph.Result()
	casecheck.True(t, len(result) == 4)
	casecheck.Contains(t, "a,b,c,d", strings.Join(result, ","))
}

func TestUnit_KahnCoherentBreakPoint2(t *testing.T) {
	graph := New()
	graph.Add("a", "b")
	graph.Add("a", "c")
	graph.Add("a", "d")
	graph.BreakPoint("w")
	casecheck.Error(t, graph.Build())
}

func TestUnit_KahnCyclical(t *testing.T) {
	graph := New()
	graph.Add("1", "2")
	graph.Add("2", "3")
	graph.Add("3", "2")
	casecheck.Error(t, graph.Build())
}

func Benchmark_Kahn1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		func() {
			graph := New()
			graph.Add("1", "2")
			graph.Add("1", "3")
			graph.Add("3", "4")
			graph.Add("2", "4")
			graph.Add("4", "5")
			_ = graph.Build() //nolint:errcheck
		}()
	}
}

func Benchmark_Kahn2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		func() {
			graph := New()
			graph.Add("1", "2")
			graph.Add("1", "3")
			graph.Add("3", "4")
			graph.Add("2", "4")
			graph.Add("4", "5")
			graph.BreakPoint("2")
			_ = graph.Build() //nolint:errcheck
		}()
	}
}
