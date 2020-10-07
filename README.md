# Algorithms
Algorithmic calculation methods

[![Release](https://img.shields.io/github/release/deweppro/go-algorithms.svg?style=flat-square)](https://github.com/deweppro/go-algorithms/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/deweppro/go-algorithms)](https://goreportcard.com/report/github.com/deweppro/go-algorithms)
[![Build Status](https://travis-ci.com/deweppro/go-algorithms.svg?branch=master)](https://travis-ci.com/deweppro/go-algorithms)

### 1) Topological sorting. Kahn's Algorithm.

```go
import (
    graphkahn "https://github.com/deweppro/go-algorithms/graph/kahn"
 )

kahn := graphkahn.New()

kahn.Add("1", "2")
kahn.Add("1", "3")
kahn.Add("3", "4")
kahn.Add("2", "4")
kahn.Add("4", "5")

kahn.Build()
result := kahn.Result()
```

**Result:** [1,2,3,4,5]