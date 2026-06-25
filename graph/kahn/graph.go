/*
 *  Copyright (c) 2019-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

// see: https://en.wikipedia.org/wiki/Topological_sorting

package kahn

import (
	"errors"
	"sort"
)

var (
	ErrBuild      = errors.New("can't do topographical sorting")
	ErrBreakPoint = errors.New("don`t found topographical break point")
)

type Graph struct {
	from map[string][]string
	to   map[string][]string

	nodes map[string]struct{}

	breakPoint string
	result     []string
}

func New() *Graph {
	return &Graph{
		from:   make(map[string][]string),
		to:     make(map[string][]string),
		nodes:  make(map[string]struct{}),
		result: make([]string, 0),
	}
}

func (g *Graph) Add(from, to string) {
	g.from[from] = append(g.from[from], to)
	g.to[to] = append(g.to[to], from)
	g.nodes[from] = struct{}{}
	g.nodes[to] = struct{}{}
}

func (g *Graph) BreakPoint(point string) {
	g.breakPoint = point
}

func (g *Graph) Build() error {
	g.result = g.result[:0]

	var active map[string]struct{}

	if len(g.breakPoint) == 0 {
		active = g.copyAllNodes()
	} else {
		active = g.copyNodesByBreakPoint()
		if len(active) == 0 {
			return ErrBreakPoint
		}
	}

	inDegree := make(map[string]int)
	for u := range active {
		for _, v := range g.from[u] {
			if _, ok := active[v]; ok {
				inDegree[v]++
			}
		}
	}

	queue := make([]string, 0, len(active))

	for _, key := range getKeys(active) {
		if inDegree[key] == 0 {
			queue = append(queue, key)
		}
	}

	for len(queue) > 0 {
		key := queue[0]
		queue = queue[1:]

		g.result = append(g.result, key)

		for _, v := range g.from[key] {
			if _, ok := active[v]; ok {
				inDegree[v]--
				if inDegree[v] == 0 {
					queue = append(queue, v)
				}
			}
		}
	}

	if len(g.result) != len(active) {
		return ErrBuild
	}

	return nil
}

func (g *Graph) Result() []string {
	return append(make([]string, 0, len(g.result)), g.result...)
}

func (g *Graph) copyAllNodes() map[string]struct{} {
	tmp := make(map[string]struct{}, len(g.nodes))
	for k := range g.nodes {
		tmp[k] = struct{}{}
	}
	return tmp
}

func (g *Graph) copyNodesByBreakPoint() map[string]struct{} {
	if _, ok := g.nodes[g.breakPoint]; !ok {
		return nil
	}

	queue := make([]string, 0, len(g.nodes))
	tmp := make(map[string]struct{}, len(g.nodes))

	queue = append(queue, g.breakPoint)
	for len(queue) > 0 {
		key := queue[0]
		queue = queue[1:]

		if _, ok := tmp[key]; ok {
			continue
		}
		tmp[key] = struct{}{}
		queue = append(queue, g.to[key]...)
	}
	return tmp
}

func getKeys(in map[string]struct{}) []string {
	result := make([]string, 0, len(in))
	for k := range in {
		result = append(result, k)
	}
	sort.Strings(result)
	return result
}
