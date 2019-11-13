# Algorithms
Algorithmic calculation methods


### 1) Topological sorting. Kahn's Algorithm.

graph/kahn/type.go:2

```go
import (
    graphkahn "https://github.com/deweppro/algorithms/graph/kahn"
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