# Package `kahn`

## Abstract

`kahn` implements Kahn's topological-sorting algorithm for string-keyed
directed graphs. An edge `from -> to` means that `from` MUST precede `to`.

## API and behavior

- `New` creates an empty graph.
- `Add(from, to)` adds both endpoint vertices and records the directed edge.
- `BreakPoint(target)` limits the next build to `target` and its transitive
  prerequisites. With no breakpoint, `Build` considers the whole graph.
- `Build` stores an ordering or returns `ErrBuild` if the active subgraph has
  a cycle. A non-empty unknown breakpoint returns `ErrBreakPoint`.
- `Result` returns a copy of the most recent successful result. Callers MUST
  call `Build` before expecting a populated result.

The graph is not synchronized; callers MUST serialize concurrent access. With
a breakpoint, edges outside the target's prerequisite subgraph are ignored.
If several valid orders exist, callers SHOULD rely only on the dependency
constraints, not on a particular tie order.

```go
package main

import (
	"fmt"
	"log"

	"go.osspkg.com/algorithms/graph/kahn"
)

func main() {
	g := kahn.New()
	g.Add("compile", "test")
	g.Add("test", "package")
	g.BreakPoint("package")
	if err := g.Build(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(g.Result())
}
```

Building the whole graph takes O(V log V + E) time: initial zero-indegree
vertices are sorted before processing. The active graph and result require
O(V + E) storage.
