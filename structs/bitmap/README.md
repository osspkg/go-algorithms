# Package `bitmap`

## Abstract

`bitmap` stores membership flags for unsigned integer indexes in a growable,
byte-backed bitmap. `Set` grows storage as needed; `Has` tests a bit and `Del`
clears it.

## API and limits

`New` accepts `OptMaxIndex` to choose initial capacity. Storage grows when
`Set` reaches a larger index. The maximum supported index is the inclusive
constant `MaxIndex` (`1 << 34`); `Set` silently ignores larger indexes, while
`Has` returns false and `Del` does nothing for them.

Internal locking is enabled by default. `OptDisableLock` removes it; callers
using that option MUST provide synchronization. Because the bounds check in
`Has` and `Del` can overlap a resize, callers MUST externally serialize a
`Set` that may grow the bitmap with concurrent reads or deletes. `CopyTo`
copies the bit storage and the source's lock setting into the destination.

`MarshalBinary` returns the raw bitmap bytes without a versioned header.
`UnmarshalBinary` replaces the receiver's contents and rejects empty or
oversized data. The serialized bytes SHOULD be stored with any metadata needed
to identify their meaning.

```go
package main

import (
	"fmt"

	"go.osspkg.com/algorithms/structs/bitmap"
)

func main() {
	bm := bitmap.New(bitmap.OptMaxIndex(1024))
	bm.Set(42)
	fmt.Println(bm.Has(42)) // true
	bm.Del(42)
}
```
