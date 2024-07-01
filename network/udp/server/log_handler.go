/*
 *  Copyright (c) 2022-2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package server

import (
	"net"

	"go.osspkg.com/x/logx"
)

type logHandler struct {
	log logx.Logger
}

func NewLogHandlerUDP(l logx.Logger) HandlerUDP {
	return &logHandler{log: l}
}

func (v *logHandler) HandlerUDP(_ Writer, addr net.Addr, b []byte) {
	v.log.WithFields(logx.Fields{
		"addr": addr.String(),
		"body": string(b),
	}).Warnf("Empty log handler UDP")
}
