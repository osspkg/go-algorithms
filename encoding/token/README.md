# Package `token`

## Abstract

`token` represents an eight-byte value as `T64`. It supports construction from
time, integers, bytes, strings, or cryptographic randomness; text and binary
marshalling; and `database/sql` scanning and value conversion.

## Representation and API

`T64.String` uses the package's process-wide encoding table and emits 18 bytes
with hyphens at positions 6 and 11. The default table is
`ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789`. `SetTable` replaces it for all
`T64` values; a table MUST contain 16–254 unique bytes, have even length, and
MUST NOT contain byte `0xff`.

- `NewByUint` stores a `uint64` in big-endian order; `Uint64` reverses it.
- `NewByTime` and `NewFormTime` store Unix nanoseconds.
- `NewByBytes` and `NewByString` store the first eight bytes of an MD5 digest.
- `NewRandom` reads a random 64-bit value. It returns `Nil` if all read
  attempts fail; the API does not return an error.
- `Parse` and `ParseBytes` return a token and a success flag.
- `Scan` accepts `string`, `[]byte`, or `nil`; empty and nil values leave the
  receiver unchanged. `Value` returns the textual form as a database string.

```go
package main

import (
	"fmt"

	"go.osspkg.com/algorithms/encoding/token"
)

func main() {
	id := token.NewByUint(42)
	text := id.String()
	parsed, ok := token.Parse(text)
	fmt.Println(text, ok, parsed == id)
}
```

## Interoperability and security considerations

`MarshalText` and `MarshalBinary` both emit the textual form. The current
`MarshalJSON` also emits that raw text without JSON quotes, so it is not a
standard JSON value; callers SHOULD encode `T64.String()` as a JSON string.
`UnmarshalJSON` expects the same raw form.

The encoding table is global: changing it changes the text form used by every
token. `NewByTime` is predictable, and the MD5-based constructors are not
collision-resistant identifiers. Applications MUST NOT use them as secrets or
authentication credentials. For untrusted parser input, callers MUST reject
byte `0xff` before calling `Parse` or `ParseBytes`; the current lookup table
does not safely index that byte.

