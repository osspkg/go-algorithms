/*
 *  Copyright (c) 2019-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package kahn_test

import (
	"errors"
	"reflect"
	"testing"

	"go.osspkg.com/algorithms/graph/kahn"
)

func Benchmark_Kahn1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		func() {
			graph := kahn.New()
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
			graph := kahn.New()
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

func TestUnit_Graph_Build(t *testing.T) {
	tests := []struct {
		name       string
		setup      func(g *kahn.Graph)
		breakPoint string
		want       []string
		wantErr    error
	}{
		{
			name: "Happy path: simple linear dependencies",
			setup: func(g *kahn.Graph) {
				g.Add("A", "B") // A -> B
				g.Add("B", "C") // B -> C
			},
			want: []string{"A", "B", "C"},
		},
		{
			name: "Happy path: multiple roots",
			setup: func(g *kahn.Graph) {
				g.Add("A", "C")
				g.Add("B", "C")
			},
			// Сортировка по ключам в getKeys гарантирует порядок A, B
			want: []string{"A", "B", "C"},
		},
		{
			name: "Cycle detection",
			setup: func(g *kahn.Graph) {
				g.Add("A", "B")
				g.Add("B", "C")
				g.Add("C", "A") // Cycle
			},
			wantErr: kahn.ErrBuild,
		},
		{
			name: "BreakPoint: subset of graph",
			setup: func(g *kahn.Graph) {
				g.Add("Base", "Lib")
				g.Add("Lib", "App")
				g.Add("Other", "Unused")
			},
			breakPoint: "App",
			want:       []string{"Base", "Lib", "App"},
		},
		{
			name: "BreakPoint: complex dependencies",
			setup: func(g *kahn.Graph) {
				g.Add("Common", "Auth")
				g.Add("Common", "DB")
				g.Add("Auth", "API")
				g.Add("DB", "API")
				g.Add("Utils", "Extra") // Должно быть проигнорировано
			},
			breakPoint: "API",
			want:       []string{"Common", "Auth", "DB", "API"},
		},
		{
			name: "BreakPoint: not found",
			setup: func(g *kahn.Graph) {
				g.Add("A", "B")
			},
			breakPoint: "Z",
			wantErr:    kahn.ErrBreakPoint,
		},
		{
			name: "Cycle outside of BreakPoint scope",
			setup: func(g *kahn.Graph) {
				g.Add("A", "B") // Path to BreakPoint
				g.Add("C", "D") // Cycle
				g.Add("D", "C") // Cycle
			},
			breakPoint: "B",
			want:       []string{"A", "B"}, // Должно работать, так как цикл не в активных узлах
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := kahn.New()
			tt.setup(g)
			if tt.breakPoint != "" {
				g.BreakPoint(tt.breakPoint)
			}

			err := g.Build()
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("Build() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Fatalf("Build() unexpected error: %v", err)
			}

			if !reflect.DeepEqual(g.Result(), tt.want) {
				t.Errorf("Result() = %v, want %v", g.Result(), tt.want)
			}
		})
	}
}

func TestUnit_Graph_ResultIsCopy(t *testing.T) {
	g := kahn.New()
	g.Add("A", "B")
	_ = g.Build()

	res1 := g.Result()
	res1[0] = "MUTATED"

	res2 := g.Result()
	if res2[0] == "MUTATED" {
		t.Error("Result() returned a pointer to the internal slice, not a copy")
	}
}
