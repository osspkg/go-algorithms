/*
 *  Copyright (c) 2019-2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package network

import (
	"net"
	"strings"

	"go.osspkg.com/x/errors"
)

var (
	ErrResolveTCPAddress = errors.New("resolve tcp address")
)

func RandomPort(host string) (string, error) {
	host = strings.Join([]string{host, "0"}, ":")
	addr, err := net.ResolveTCPAddr("tcp", host)
	if err != nil {
		return host, errors.Wrap(err, ErrResolveTCPAddress)
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return host, errors.Wrap(err, ErrResolveTCPAddress)
	}
	v := l.Addr().String()
	if err = l.Close(); err != nil {
		return host, errors.Wrap(err, ErrResolveTCPAddress)
	}
	return v, nil
}

func CheckHostPort(addr string) string {
	hp := strings.Split(addr, ":")
	if len(hp) != 2 {
		tmp := make([]string, 2)
		copy(hp, tmp)
		hp = tmp
	}
	if len(hp[0]) == 0 {
		hp[0] = "0.0.0.0"
	}
	if len(hp[1]) == 0 {
		if v, err := RandomPort(hp[0]); err == nil {
			hp[1] = v
		} else {
			hp[1] = "8080"
		}
	}
	return strings.Join(hp, ":")
}

func Normalize(p string, ips ...string) []string {
	result := make([]string, 0, len(ips))
	for _, ip := range ips {
		host, port, err := net.SplitHostPort(ip)
		if err != nil {
			host = ip
			port = p
		}
		if !IsValidIP(host) {
			continue
		}
		if port == "0" {
			port = p
		}
		result = append(result, net.JoinHostPort(host, port))
	}

	return result
}

func IsValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}
