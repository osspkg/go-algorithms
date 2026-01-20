/*
 *  Copyright (c) 2019-2026 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package otp

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

type Option func(o *OTP)

func OptHashSHA1() Option {
	return func(o *OTP) {
		o.hash = sha1.New
		o.algorithm = "SHA1"
	}
}

func OptHashSHA256() Option {
	return func(o *OTP) {
		o.hash = sha256.New
		o.algorithm = "SHA256"
	}
}

func OptHashSHA512() Option {
	return func(o *OTP) {
		o.hash = sha512.New
		o.algorithm = "SHA512"
	}
}

func OptPeriod(v int64) Option {
	return func(o *OTP) {
		if v < 30 {
			v = 30
		}
		if v > 120 {
			v = 120
		}
		o.period = v
	}
}

func OptCode6Digits() Option {
	return func(o *OTP) {
		o.codeSize = 6
		o.codeTmpl = fmt.Sprintf("%%0%dd", o.codeSize)
	}
}

func OptCode8Digits() Option {
	return func(o *OTP) {
		o.codeSize = 8
		o.codeTmpl = fmt.Sprintf("%%0%dd", o.codeSize)
	}
}
