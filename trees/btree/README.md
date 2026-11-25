# Package `btree`

## Abstract

`btree` provides a concurrent generic B-tree for keys ordered by Go's
`cmp.Ordered` constraint and arbitrary values.

## API and behavior

- `New[K, V](degree)` creates a tree. Degrees below two are raised to two.
- `Insert(key, value)` inserts a key or replaces the value for an existing key.
- `Find(key)` returns the value and `true` when present; otherwise it returns
  the zero value of `V` and `false`.
- `Delete(key)` removes a key if present and otherwise does nothing.

The tree uses internal synchronization by default and supports concurrent
operations. Keys MUST have a consistent total ordering. In particular, callers
MUST NOT insert floating-point NaN keys because ordinary comparisons do not
order NaN consistently.

```go
package main

import (
	"fmt"

	"go.osspkg.com/algorithms/trees/btree"
)

func main() {
	tree := btree.New[int, string](16)
	tree.Insert(42, "answer")
	if value, ok := tree.Find(42); ok {
		fmt.Println(value)
	}
	tree.Delete(42)
}
```

The degree controls the number of keys stored per node and SHOULD be selected
for the expected key type, workload, and memory/cache trade-offs. Operations
take logarithmic tree depth with work inside each node proportional to its key
count.
