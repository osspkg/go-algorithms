/*
 *  Copyright (c) 2019-2024 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

// see: https://en.wikipedia.org/wiki/Topological_sorting

package kahn

import (
	"errors"
	"fmt"
	"sort"
)

var (
	ErrBuildKahn      = errors.New("can't do topographical sorting")
	ErrBreakPointKahn = errors.New("don`t found topographical break point")
)

const empty = ""

type Graph struct {
	graph      map[string]map[string]int
	all        map[string]struct{}
	result     []string
	breakPoint string
}

func New() *Graph {
	return &Graph{
		graph:  make(map[string]map[string]int),
		all:    make(map[string]struct{}),
		result: make([]string, 0),
	}
}

// Add - Adding a graph edge
func (k *Graph) Add(from, to string) {
	if _, ok := k.graph[from]; !ok {
		k.graph[from] = make(map[string]int)
	}
	k.graph[from][to]++
}

func (k *Graph) BreakPoint(point string) {
	k.breakPoint = point
}

// To update the temporary map
func (k *Graph) updateTemp() (int, []string) {
	for i, sub := range k.graph {
		for j := range sub {
			k.all[j] = struct{}{}
		}
		k.all[i] = struct{}{}
	}
	temp := make([]string, 0, len(k.all))
	for s := range k.all {
		temp = append(temp, s)
	}
	sort.Strings(temp)
	return len(k.all), temp
}

// Build - Perform sorting
func (k *Graph) Build() error {
	k.result = k.result[:0]
	length, temp := k.updateTemp()

	if len(k.breakPoint) > 0 {
		j := -1
		for i, name := range temp {
			if k.breakPoint == name {
				j = i
			}
		}
		if j < 0 {
			return fmt.Errorf("%w: %s", ErrBreakPointKahn, k.breakPoint)
		}
		temp[0], temp[j] = temp[j], temp[0]
	}

	for len(k.result) < length {
		found := ""
		i := 0
		for j, item := range temp {
			if item == empty {
				continue
			}
			if k.find(item) {
				found = item
				i = j
				break
			}
		}
		if len(found) > 0 {
			k.result = append(k.result, found)
			delete(k.all, found)
			temp[i] = empty
		} else {
			return ErrBuildKahn
		}
		if len(k.breakPoint) > 0 && found == k.breakPoint {
			break
		}
	}
	return nil
}

// Finding the next edge
func (k *Graph) find(item string) bool {
	for i, j := range k.graph {
		if _, jok := j[item]; jok {
			if _, iok := k.all[i]; iok {
				return false
			}
		}
	}
	return true
}

// Result - Getting a sorted slice
func (k *Graph) Result() []string {
	return append(make([]string, 0, len(k.result)), k.result...)
}
