# Package `comp`

## Abstract

`comp` compares sequences of comparable values. Its constructors return a
`Comparer[T]`, whose public interface exposes a normalized `Similarity` score.
Use `[]rune` for Unicode text, `[]byte` for byte sequences, or slices of other
comparable values for tokens and symbols.

## Algorithms

| Constructor | Measures |
| --- | --- |
| `NewHamming[T]` | Positional differences; input lengths must match for a non-zero similarity. |
| `NewLevenshtein[T]` | Insertions, deletions, and substitutions. |
| `NewDamerauLevenshtein[T]` | Levenshtein edits plus adjacent transpositions using the restricted optimal-string-alignment variant. |
| `NewLCS[T]` | Longest common subsequence, not longest common substring. |
| `NewJaroWinkler[T]` | Jaro-style matching-window similarity. The current implementation does not apply a separate Winkler prefix boost. |

`Similarity` returns a score from `0` to `100`, rounded to two decimal
places. Two empty sequences score `100`. Hamming comparison of unequal-length
sequences scores `0`.

```go
package main

import (
	"fmt"

	"go.osspkg.com/algorithms/comp"
)

func main() {
	comparer := comp.NewLevenshtein[rune]()
	score := comparer.Similarity([]rune("kitten"), []rune("sitting"))
	fmt.Printf("%.2f\n", score)
}
```

The constructors expose `Similarity` through `Comparer[T]`; distance methods
are not part of that interface.
