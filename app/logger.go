/*
 *  Copyright (c) 2019-2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package app

import (
	"os"

	"go.osspkg.com/x/log"
)

type _log struct {
	file    *os.File
	handler log.Logger
	conf    LogConfig
}

func newLog(conf LogConfig) *_log {
	file, err := os.OpenFile(conf.FilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return &_log{file: file, conf: conf}
}

func (v *_log) Handler(l log.Logger) {
	v.handler = l
	v.handler.SetOutput(v.file)
	v.handler.SetLevel(v.conf.Level)
	switch v.conf.Format {
	case "string":
		v.handler.SetFormatter(log.NewFormatString())
	case "json":
		v.handler.SetFormatter(log.NewFormatJSON())
	}
}

func (v *_log) Close() error {
	if v.handler != nil {
		v.handler.Close()
	}
	return v.file.Close()
}
