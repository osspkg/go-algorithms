/*
 *  Copyright (c) 2019-2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package logx

import "io"

var std = New()

// Default logger
func Default() Logger {
	return std
}

// SetOutput change writer
func SetOutput(out io.Writer) {
	std.SetOutput(out)
}

func SetFormatter(f Formatter) {
	std.SetFormatter(f)
}

// SetLevel change log level
func SetLevel(v uint32) {
	std.SetLevel(v)
}

// GetLevel getting log level
func GetLevel() uint32 {
	return std.GetLevel()
}

// Close waiting for all messages to finish recording
func Close() {
	std.Close()
}

// Infof info message
func Infof(format string, args ...interface{}) {
	std.Infof(format, args...)
}

// Warnf warning message
func Warnf(format string, args ...interface{}) {
	std.Warnf(format, args...)
}

// Errorf error message
func Errorf(format string, args ...interface{}) {
	std.Errorf(format, args...)
}

// Debugf debug message
func Debugf(format string, args ...interface{}) {
	std.Debugf(format, args...)
}

// Fatalf fatal message and exit
func Fatalf(format string, args ...interface{}) {
	std.Fatalf(format, args...)
}

// WithFields setter context to log message
func WithFields(v Fields) Writer {
	return std.WithFields(v)
}

// WithError setter context to log message
func WithError(key string, err error) Writer {
	return std.WithError(key, err)
}

// WithField setter context to log message
func WithField(key string, value interface{}) Writer {
	return std.WithField(key, value)
}
