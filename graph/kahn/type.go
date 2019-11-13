// see: https://en.wikipedia.org/wiki/Topological_sorting
package kahn

import (
	"errors"
)

// Kahn ___
type Kahn struct {
	graph  map[string]map[string]int
	tmp    map[string]bool
	result []string
}

// New ___
func New() *Kahn {
	return &Kahn{
		graph:  make(map[string]map[string]int),
		tmp:    make(map[string]bool),
		result: make([]string, 0),
	}
}

var (
	errBuild = errors.New("Can't do topographical sorting")
)

// Add - Adding a graph edge
func (k *Kahn) Add(from, to string) error {
	if _, ok := k.graph[from]; !ok {
		k.graph[from] = make(map[string]int)
	}

	k.graph[from][to]++

	return nil
}

// To update the temporary map
func (k *Kahn) updateTemp() int {
	for i, sub := range k.graph {
		for j, _ := range sub {
			k.tmp[j] = true
		}
		k.tmp[i] = true
	}

	return len(k.tmp)
}

// Build - Perform sorting
func (k *Kahn) Build() error {
	k.result = k.result[:0]
	length := k.updateTemp()

	for len(k.result) < length {
		found := ""

		for item, _ := range k.tmp {
			if k.find(item) {
				found = item
				break
			}
		}

		if len(found) > 0 {
			k.result = append(k.result, found)
			delete(k.tmp, found)
		} else {
			return errBuild
		}
	}

	return nil
}

// Finding the next edge
func (k *Kahn) find(item string) bool {
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
func (k *Kahn) Result() []string {
	return k.result
}
