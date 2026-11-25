# Package `sorts`

## Abstract

`sorts` provides generic in-place sorting algorithms and slice reversal. Each
sort accepts a slice and a comparison callback that receives indexes into the
slice's current, partially sorted state.

## Operations

| Function | Time | Extra storage | Stable |
| --- | --- | --- | --- |
| `Bubble` | O(n²), O(n) when already sorted | O(1) | Yes |
| `Cocktail` | O(n²) | O(1) | Yes |
| `Insertion` | O(n²), O(n) when already sorted | O(1) | Yes |
| `Selection` | O(n²) | O(1) | No |
| `Heapsort` | O(n log n) | O(1) | No |
| `Merge` | O(n log n) | O(n) | No guarantee |
| `Reverse` | O(n) | O(1) | N/A |

`less(i, j)` MUST return true when the element currently at index `i` should
precede the element at index `j`. It SHOULD define a strict weak ordering.
Because sorting mutates the slice, the callback MUST read the current indexed
elements rather than assume that indexes still refer to their original values.

```go
package main

import (
	"fmt"

	"go.osspkg.com/algorithms/sorts"
)

func main() {
	values := []int{5, 1, 4, 2, 3}
	sorts.Merge(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	fmt.Println(values) // [1 2 3 4 5]
}
```

All sorting functions mutate the supplied slice. `Merge` allocates a temporary
buffer; the other operations use constant auxiliary storage. For general
application sorting, callers SHOULD also consider `slices.SortFunc` or
`sort.Slice` from the Go standard library.
