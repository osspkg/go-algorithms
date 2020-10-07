/*
 * Copyright (c) 2020.  Mikhail Knyazhev <markus621@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/gpl-3.0.html>.
 */
// see: https://en.wikipedia.org/wiki/Topological_sorting

package kahn

import "github.com/pkg/errors"

var (
	ErrBuildKahn = errors.New("can't do topographical sorting")
)

type Graph struct {
	graph  map[string]map[string]int
	tmp    map[string]bool
	result []string
}

func New() *Graph {
	return &Graph{
		graph:  make(map[string]map[string]int),
		tmp:    make(map[string]bool),
		result: make([]string, 0),
	}
}

// Add - Adding a graph edge
func (k *Graph) Add(from, to string) error {
	if _, ok := k.graph[from]; !ok {
		k.graph[from] = make(map[string]int)
	}
	k.graph[from][to]++
	return nil
}

// To update the temporary map
func (k *Graph) updateTemp() int {
	for i, sub := range k.graph {
		for j := range sub {
			k.tmp[j] = true
		}
		k.tmp[i] = true
	}
	return len(k.tmp)
}

// Build - Perform sorting
func (k *Graph) Build() error {
	k.result = k.result[:0]
	length := k.updateTemp()
	for len(k.result) < length {
		found := ""
		for item := range k.tmp {
			if k.find(item) {
				found = item
				break
			}
		}
		if len(found) > 0 {
			k.result = append(k.result, found)
			delete(k.tmp, found)
		} else {
			return ErrBuildKahn
		}
	}
	return nil
}

// Finding the next edge
func (k *Graph) find(item string) bool {
	for i, j := range k.graph {
		if _, jok := j[item]; jok {
			if _, iok := k.tmp[i]; iok {
				return false
			}
		}
	}
	return true
}

// Result - Getting a sorted slice
func (k *Graph) Result() []string {
	return k.result
}
