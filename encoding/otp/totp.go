/*
 *  Copyright (c) 2019-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package otp

import (
	"crypto/hmac"
	"crypto/rand"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"hash"
	"io"
	"math"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var b32 = base32.StdEncoding.WithPadding(base32.NoPadding)

type OTP struct {
	hash      func() hash.Hash
	algorithm string

	period int64

	codeSize int
	codeTmpl string
}

func New(options ...Option) (*OTP, error) {
	obj := &OTP{}

	opts := make([]Option, 0, 10)
	opts = append(opts, OptHashSHA1(), OptPeriod(30), OptCode6Digits())
	opts = append(opts, options...)

	for _, opt := range opts {
		opt(obj)
	}

	return obj, nil
}

func (o *OTP) NewSecret(size int) (string, error) {
	if size < 1 {
		size = 10
	}
	secret := make([]byte, size)
	if _, err := io.ReadFull(rand.Reader, secret); err != nil {
		return "", err
	}
	return b32.EncodeToString(secret), nil
}

func (o *OTP) validateSecret(secret string) ([]byte, error) {
	secret = strings.TrimSpace(secret)
	if n := len(secret) % 8; n != 0 {
		secret = secret + strings.Repeat("=", 8-n)
	}
	secret = strings.ToUpper(secret)
	secretBytes, err := base32.StdEncoding.DecodeString(secret)
	if err != nil {
		return nil, fmt.Errorf("invalid secret: %w", err)
	}
	return secretBytes, nil
}

func (o *OTP) generate(secret string, counter uint64, delta int64) (string, error) {
	b, err := o.validateSecret(secret)
	if err != nil {
		return "", err
	}

	counterBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(counterBytes, uint64(int64(counter)+delta))

	hm := hmac.New(o.hash, b)
	if _, err := hm.Write(counterBytes); err != nil {
		return strings.Repeat("0", o.codeSize), err
	}
	timeHash := hm.Sum(nil)

	offset := int(timeHash[len(timeHash)-1] & 0x0F)
	truncHash := int64(
		(int(timeHash[offset])&0x7f)<<24 |
			(int(timeHash[offset+1])&0xff)<<16 |
			(int(timeHash[offset+2])&0xff)<<8 |
			(int(timeHash[offset+3]) & 0xff))

	otp := truncHash % int64(math.Pow10(o.codeSize))

	return fmt.Sprintf(o.codeTmpl, otp), nil
}

func (o *OTP) GenerateTOTP(secret string, delta int64) (string, error) {
	currentTime := time.Now().Unix()
	counter := uint64(math.Floor(float64(currentTime) / float64(o.period)))

	return o.generate(secret, counter, delta)
}

func (o *OTP) GenerateHOTP(secret string, counter uint64) (string, error) {
	return o.generate(secret, counter, 0)
}

func (o *OTP) UrlTOTP(secret, account, issuer string) string {
	secret = strings.TrimSpace(secret)
	params := url.Values{
		"secret":    []string{secret},
		"issuer":    []string{issuer},
		"algorithm": []string{o.algorithm},
		"digits":    []string{strconv.Itoa(o.codeSize)},
		"period":    []string{strconv.Itoa(int(o.period))},
	}

	uri := url.URL{
		Scheme:   "otpauth",
		Host:     "totp",
		Path:     "/" + account,
		RawQuery: params.Encode(),
	}

	return uri.String()
}

func (o *OTP) UrlHOTP(secret, account, issuer string, counter uint64) string {
	secret = strings.TrimSpace(secret)
	params := url.Values{
		"secret":    []string{secret},
		"issuer":    []string{issuer},
		"algorithm": []string{o.algorithm},
		"digits":    []string{strconv.Itoa(o.codeSize)},
		"period":    []string{strconv.Itoa(int(o.period))},
		"counter":   []string{strconv.Itoa(int(counter))},
	}

	uri := url.URL{
		Scheme:   "otpauth",
		Host:     "hotp",
		Path:     "/" + account,
		RawQuery: params.Encode(),
	}

	return uri.String()
}
