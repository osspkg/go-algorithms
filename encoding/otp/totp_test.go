/*
 *  Copyright (c) 2019-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package otp_test

import (
	"testing"

	"go.osspkg.com/casecheck"

	"go.osspkg.com/algorithms/encoding/otp"
)

func TestUnit_TOTP_Generate(t *testing.T) {
	gen, err := otp.New()
	casecheck.NoError(t, err)

	c1, err := gen.GenerateTOTP(`4QEXNRSWEYM5HWCG`, 0)
	casecheck.NoError(t, err)
	t.Log(c1)
	c2, err := gen.GenerateTOTP(`4QEXNRSWEYM5HWCG`, 0)
	casecheck.NoError(t, err)
	c3, err := gen.GenerateTOTP(`JBSWY3DPEHPK3PXP`, 0)
	casecheck.NoError(t, err)

	casecheck.Equal(t, c1, c2)
	casecheck.NotEqual(t, c1, c3)

	link := gen.UrlTOTP(`JBSWY3DPEHPK3PXP`, `user name`, `example.com`)
	want := `otpauth://totp/user%20name?algorithm=SHA1&digits=6&issuer=example.com&period=30&secret=JBSWY3DPEHPK3PXP`
	casecheck.Equal(t, want, link)
}

func TestUnit_HOTP_Generate(t *testing.T) {
	gen, err := otp.New()
	casecheck.NoError(t, err)

	c1, err := gen.GenerateHOTP(`4QEXNRSWEYM5HWCG`, 0)
	casecheck.NoError(t, err)
	c2, err := gen.GenerateHOTP(`4QEXNRSWEYM5HWCG`, 0)
	casecheck.NoError(t, err)
	c3, err := gen.GenerateHOTP(`4QEXNRSWEYM5HWCG`, 1)
	casecheck.NoError(t, err)
	c4, err := gen.GenerateHOTP(`JBSWY3DPEHPK3PXP`, 0)
	casecheck.NoError(t, err)

	casecheck.Equal(t, c1, c2)
	casecheck.NotEqual(t, c1, c3)
	casecheck.NotEqual(t, c1, c4)

	link := gen.UrlHOTP(`JBSWY3DPEHPK3PXP`, `user name`, `example.com`, 0)
	want := `otpauth://hotp/user%20name?algorithm=SHA1&counter=0&digits=6&issuer=example.com&period=30&secret=JBSWY3DPEHPK3PXP`
	casecheck.Equal(t, want, link)
}

/*
goos: linux
goarch: amd64
pkg: go.osspkg.com/algorithms/encoding/totp
cpu: 12th Gen Intel(R) Core(TM) i9-12900KF
Benchmark_TOTP_Generate
Benchmark_TOTP_Generate-4   	 6530068	       180.6 ns/op	     512 B/op	      10 allocs/op
*/
func Benchmark_TOTP_Generate(b *testing.B) {
	gen, err := otp.New()
	casecheck.NoError(b, err)

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			gen.GenerateTOTP(`4QEXNRSWEYM5HWCG`, 0)
		}
	})
}
