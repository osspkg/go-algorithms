# Algorithms

Algorithmic calculation methods

[![Coverage Status](https://coveralls.io/repos/github/osspkg/go-algorithms/badge.svg?branch=master)](https://coveralls.io/github/osspkg/go-algorithms?branch=master)
[![Release](https://img.shields.io/github/release/osspkg/go-algorithms.svg?style=flat-square)](https://github.com/osspkg/go-algorithms/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/osspkg/go-algorithms)](https://goreportcard.com/report/github.com/osspkg/go-algorithms)
[![CI](https://github.com/osspkg/go-algorithms/actions/workflows/ci.yml/badge.svg)](https://github.com/osspkg/go-algorithms/actions/workflows/ci.yml)

## Install

```shell
go get -u go.osspkg.com/algorithms
```

## List

- Graph
  - Topological sorting
    - [Kahn's Algorithm](graph/kahn/type.go)
- Information compression
  - [Reducing numbers](shorten/shorten.go)
- Sorting algorithm
  - [Bubble sort](sorts/bubble.go)
  - [Cocktail shaker sort](sorts/cocktail.go)
  - [Insertion sort](sorts/insertion.go)
  - [Merge sort](sorts/merge.go)
  - [Selection sort](sorts/selection.go)
  - [Heapsort](sorts/heapsort.go)
- Filtering algorithms
  - [Bloom filter](filters/bloom/bloom.go)

## License

BSD-3-Clause License. See the LICENSE file for details
