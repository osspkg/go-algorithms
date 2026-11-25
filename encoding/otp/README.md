# Package `otp`

## Abstract

`otp` generates HOTP and TOTP values using HMAC and creates `otpauth`-style
provisioning URIs. HOTP follows the counter-based construction in [RFC 4226];
TOTP uses the time-based moving factor described in [RFC 6238]. This package
generates values; it does not verify them, enforce a replay policy, or choose
a verifier's accepted clock window.

## Configuration

`New` defaults to HMAC-SHA1, a 30-second TOTP period, and six digits. Options
can select SHA-1, SHA-256, or SHA-512; six or eight digits; and a period. The
period is clamped to the inclusive range 30–120 seconds.

`NewSecret(size)` generates `size` random bytes and returns unpadded Base32.
Values below one use a 10-byte default. `GenerateHOTP` takes an explicit
counter. `GenerateTOTP` uses the current Unix time; its `delta` argument shifts
the time-step counter, not seconds.

`UrlTOTP` and `UrlHOTP` create provisioning URIs using the `otpauth` scheme.
`UrlHOTP` includes both the counter and the configured period in its query.

```go
package main

import (
	"fmt"
	"log"

	"go.osspkg.com/algorithms/encoding/otp"
)

func main() {
	generator, err := otp.New(otp.OptHashSHA256(), otp.OptCode8Digits())
	if err != nil {
		log.Fatal(err)
	}
	secret, err := generator.NewSecret(20)
	if err != nil {
		log.Fatal(err)
	}
	code, err := generator.GenerateTOTP(secret, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(code)
}
```

The URI methods encode account and issuer values as URL components. Provisioning
URIs contain the secret and MUST be treated as credentials. Applications
SHOULD protect secrets at rest and in transit, use synchronized clocks, and
avoid logging codes, secrets, or generated URIs.
