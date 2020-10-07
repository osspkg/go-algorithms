# Algorithms
Algorithmic calculation methods

[![Release](https://img.shields.io/github/release/deweppro/go-algorithms.svg?style=flat-square)](https://github.com/deweppro/go-algorithms/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/deweppro/go-algorithms)](https://goreportcard.com/report/github.com/deweppro/go-algorithms)
[![Build Status](https://travis-ci.com/deweppro/go-algorithms.svg?branch=master)](https://travis-ci.com/deweppro/go-algorithms)

### 1) Topological sorting. Kahn's Algorithm.

```go
import (
    "https://github.com/deweppro/go-algorithms/graph/kahn"
 )

graph := kahn.New()

graph.Add("1", "2")
graph.Add("1", "3")
graph.Add("3", "4")
graph.Add("2", "4")
graph.Add("4", "5")

graph.Build()
result := graph.Result()
```

**Result:** [1,2,3,4,5]