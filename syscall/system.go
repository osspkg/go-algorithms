/*
 *  Copyright (c) 2019-2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package syscall

import (
	"os"
	"os/signal"
	"strconv"
	scall "syscall"
)

// OnStop calling a function if you send a system event stop
func OnStop(callFunc func()) {
	quit := make(chan os.Signal, 4)
	signal.Notify(quit, os.Interrupt, scall.SIGINT, scall.SIGTERM, scall.SIGKILL) //nolint:staticcheck
	<-quit
	signal.Stop(quit)
	callFunc()
}

// OnUp calling a function if you send a system event SIGHUP
func OnUp(callFunc func()) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, scall.SIGHUP)
	<-quit
	signal.Stop(quit)
	callFunc()
}

// OnCustom calling a function if you send a system custom event
func OnCustom(callFunc func(), sig ...os.Signal) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, sig...)
	<-quit
	signal.Stop(quit)
	callFunc()
}

// Pid write pid file
func Pid(filename string) error {
	pid := strconv.Itoa(scall.Getpid())
	return os.WriteFile(filename, []byte(pid), 0755)
}
