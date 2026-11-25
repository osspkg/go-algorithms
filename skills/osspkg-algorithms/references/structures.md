# Data structures

## Bitmap: `structs/bitmap`

```go
package main

import (
	"log"

	"go.osspkg.com/algorithms/structs/bitmap"
)

func main() {
	bm := bitmap.New(bitmap.OptMaxIndex(1024))
	bm.Set(42)
	if !bm.Has(42) {
		log.Fatal("bit was not set")
	}
	bm.Del(42)

	data, err := bm.MarshalBinary()
	if err != nil {
		log.Fatal(err)
	}
	copyBM := bitmap.New()
	if err := copyBM.UnmarshalBinary(data); err != nil {
		log.Fatal(err)
	}
}
```

`Set`, `Del`, and `Has` use `uint64` indices; the bitmap grows automatically to
the requested index, but values above the exported `bitmap.MaxIndex` are
ignored or treated as absent. `MarshalBinary` returns a copy of the raw bytes.
`UnmarshalBinary` rejects empty input. The raw format has no versioned header.

Operations are protected by an `RWMutex` by default. Use `OptDisableLock()` only
when access is strictly single-threaded or synchronization is fully handled
externally; it is not a safe default optimization. Even with the default lock,
externally serialize a `Set` that may grow storage with concurrent `Has` or
`Del` calls. `CopyTo` copies the source's lock setting along with its bit data.

## Bloom filter: `structs/bloom`

```go
package main

import (
	"bytes"
	"log"

	"go.osspkg.com/algorithms/structs/bloom"
)

func main() {
	filter, err := bloom.New(bloom.Quantity(100_000, 0.01))
	if err != nil {
		log.Fatal(err)
	}
	filter.Add("user:42")
	if !filter.Contain("user:42") {
		log.Fatal("inserted value is absent")
	}

	var buf bytes.Buffer
	if err := filter.Dump(&buf); err != nil {
		log.Fatal(err)
	}

	restored, err := bloom.New(bloom.Quantity(100_000, 0.01))
	if err != nil {
		log.Fatal(err)
	}
	if err := restored.Restore(bytes.NewReader(buf.Bytes())); err != nil {
		log.Fatal(err)
	}
}
```

`Quantity(expectedItems, falsePositiveRate)` sets the expected number of items
and the target false-positive probability; the size must be greater than zero,
and the rate must be strictly between `0` and `1`. Defaults are 10,000,000
items and a 0.1 false-positive rate; xxhash is used by default. With `HashFunc`,
provide a non-nil `func() hash.Hash` factory such as `sha256.New`. Each filter
gets random salts, so independently created filters are not interchangeable.

The filter answers "definitely not" or "possibly present": false positives are
normal, and `Contain` does not replace the source of truth. Use the same key
representation with `Add` and `Contain`; the library converts strings, bytes,
integers, and values implementing supported interfaces into bytes.

`Dump` stores the salts and bitmap, but not the hash function or the original
`Quantity`; `Restore` is limited to 64 MiB. To restore, create the destination
with the same quantity configuration and hash factory as the source, then
check the error from `Restore`. `CopyTo` copies bits and salts but does not
replace the destination's hash factory, which MUST match the source.

## B-tree: `trees/btree`

```go
package main

import (
	"fmt"

	"go.osspkg.com/algorithms/trees/btree"
)

func main() {
	tree := btree.New[int, string](2)
	tree.Insert(10, "ten")
	tree.Insert(20, "twenty")
	tree.Insert(10, "TEN") // duplicate key updates the value

	if value, ok := tree.Find(10); ok {
		fmt.Println(value) // TEN
	}
	tree.Delete(20)
}
```

Key `K` must satisfy `cmp.Ordered`, while value `V` can be any type. `degree < 2`
is automatically replaced with `2`. `Insert` updates the value for a duplicate
key; `Find` returns `(zeroValue, false)` for a missing key, so always check the
`bool`. `Delete` does nothing for a missing key. There is no public traversal or
iterator.

`BTree` protects `Find`, `Insert`, and `Delete` with an internal `RWMutex`, but
do not copy the value after creation: pass around the pointer returned by `New`.
Keys must have a consistent total ordering; do not use floating-point NaN keys.
The degree controls keys per node, so choose it for the key type, workload, and
memory/cache trade-offs.

Verified examples in the checkout: `structs/bitmap/bitmap_test.go`,
`structs/bloom/bloom_test.go`, and `trees/btree/btree_test.go`.
