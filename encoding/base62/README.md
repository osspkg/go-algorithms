# Package `base62`

## Abstract

`base62` encodes and decodes `uint64` values using a caller-supplied alphabet
of 62 unique bytes. The alphabet defines the digit order and is kept by the
`Base62` value for subsequent operations.

## API and behavior

`New(alphabet)` panics unless the alphabet is exactly 62 bytes long and every
byte is unique. Callers MUST validate externally supplied alphabets before
constructing a codec. `Encode(0)` returns the empty string. `Decode` returns
zero for empty input, invalid characters, or arithmetic overflow; its return
value does not distinguish those cases from a valid encoding of zero.

```go
package main

import (
	"fmt"

	"go.osspkg.com/algorithms/encoding/base62"
)

func main() {
	codec := base62.New("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	text := codec.Encode(123456)
	value := codec.Decode(text)
	fmt.Println(text, value)
}
```

Base62 is a representation, not encryption or authentication. Callers SHOULD
not use an encoded value to hide or protect sensitive data.
