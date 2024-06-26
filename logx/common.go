/*
 *  Copyright (c) 2019-2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package logx

import "io"

const (
	levelFatal uint32 = iota
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
)

var levels = map[uint32]string{
	levelFatal: "FAT",
	LevelError: "ERR",
	LevelWarn:  "WRN",
	LevelInfo:  "INF",
	LevelDebug: "DBG",
}

type Fields map[string]interface{}

type Sender interface {
	PutEntity(v *entity)
	SendMessage(level uint32, call func(v *Message))
	Close()
}

// Writer interface
type Writer interface {
	Fatalf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
}

type WriterContext interface {
	WithError(key string, err error) Writer
	WithField(key string, value interface{}) Writer
	WithFields(Fields) Writer
	Writer
}

// Logger base interface
type Logger interface {
	SetOutput(out io.Writer)
	SetFormatter(f Formatter)
	SetLevel(v uint32)
	GetLevel() uint32
	Close()

	WriterContext
}
