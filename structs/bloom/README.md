# Package `bloom`

## Abstract

`bloom` implements a space-efficient probabilistic membership filter. A
negative `Contain` result means the value is definitely absent; a positive
result means it may be present. False positives are possible; false negatives
are not expected when the filter is used and restored with matching data.

## Configuration and operations

`New` defaults to an expected capacity of 10,000,000 values and a target false
positive rate of 0.1. `Quantity(n, p)` sets those values; `n` MUST be non-zero
and `p` MUST be strictly between zero and one. `HashFunc` selects a non-nil
`hash.Hash` factory; xxhash is the default. Each new filter also gets random
salts, so independently created filters are not interchangeable.

`Add` and `Contain` accept strings, byte slices, integer types, and values
implementing supported byte, string, or encoding interfaces. Other values are
formatted with `fmt`. Callers SHOULD use a stable byte representation for a
value across additions, lookups, and persistence.

```go
package main

import (
	"fmt"
	"log"

	"go.osspkg.com/algorithms/structs/bloom"
)

func main() {
	filter, err := bloom.New(bloom.Quantity(100_000, 0.01))
	if err != nil {
		log.Fatal(err)
	}
	filter.Add("user:42")
	fmt.Println(filter.Contain("user:42")) // true
}
```

The filter is synchronized for concurrent use. `Dump` and `Restore` use a
package-specific binary format; `Restore` is limited to 64 MiB. The dump does
not contain the sizing parameters, so callers MUST create the destination
with matching `Quantity` and `HashFunc` settings before restoring. `CopyTo`
copies the bits and salts but does not replace the destination's hash factory,
so the destination MUST use the same hash configuration. A Bloom filter is
not a source of truth and MUST NOT be used as an authorization decision by
itself.
