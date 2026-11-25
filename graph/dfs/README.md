# Package `dfs`

## Abstract

`dfs` builds generic directed graphs and returns a depth-first topological
ordering. An edge `from -> to` means `from` MUST precede `to` in the result.

## API and behavior

- `NewGraph[K comparable]` creates an empty, synchronized graph.
- `AddNode` registers a vertex. Adding an existing key returns an error that
  wraps `ErrNodeKeyExist`.
- `AddEdge` requires both vertices to exist and returns an error wrapping
  `ErrNodeNotFound` otherwise. Repeated edges are stored once.
- `TopologicalSort` returns an ordering for a DAG or `ErrCycleDetected` when a
  cycle is reachable in the graph.

Callers SHOULD add vertices before edges and SHOULD test errors with
`errors.Is`. The order of unrelated vertices is not guaranteed because map
iteration determines traversal order.

```go
package main

import (
	"fmt"
	"log"

	"go.osspkg.com/algorithms/graph/dfs"
)

func main() {
	g := dfs.NewGraph[string]()
	for _, node := range []string{"compile", "test"} {
		if err := g.AddNode(node); err != nil {
			log.Fatal(err)
		}
	}
	if err := g.AddEdge("compile", "test"); err != nil {
		log.Fatal(err)
	}
	order, err := g.TopologicalSort()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(order)
}
```

The traversal takes O(V + E) time and uses O(V) auxiliary state. The graph
protects mutations and sorting with an internal read-write mutex.
