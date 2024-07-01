/*
 *  Copyright (c) 2019-2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package log

import (
	"fmt"
	"io"
	"os"
	"testing"
	"time"

	"go.osspkg.com/x/sync"
	"go.osspkg.com/x/test"
)

func TestUnit_NewJSON(t *testing.T) {
	test.NotNil(t, Default())

	filename, err := os.CreateTemp(os.TempDir(), "test_new_default-*.log")
	test.NoError(t, err)

	SetOutput(filename)
	SetLevel(LevelDebug)
	test.Equal(t, LevelDebug, GetLevel())

	go Infof("async %d", 1)
	go Warnf("async %d", 2)
	go Errorf("async %d", 3)
	go Debugf("async %d", 4)

	Infof("sync %d", 1)
	Warnf("sync %d", 2)
	Errorf("sync %d", 3)
	Debugf("sync %d", 4)

	WithFields(Fields{"ip": "0.0.0.0"}).Infof("context1")
	WithFields(Fields{"nil": nil}).Infof("context2")
	WithFields(Fields{"func": func() {}}).Infof("context3")

	WithField("ip", "0.0.0.0").Infof("context4")
	WithField("nil", nil).Infof("context5")
	WithField("func", func() {}).Infof("context6")

	WithError("err", nil).Infof("context7")
	WithError("err", fmt.Errorf("er1")).Infof("context8")

	<-time.After(time.Second * 1)
	Close()

	test.NoError(t, filename.Close())
	data, err := os.ReadFile(filename.Name())
	test.NoError(t, err)
	test.NoError(t, os.Remove(filename.Name()))

	sdata := string(data)
	test.Contains(t, sdata, `"lvl":"INF","msg":"async 1"`)
	test.Contains(t, sdata, `"lvl":"WRN","msg":"async 2"`)
	test.Contains(t, sdata, `"lvl":"ERR","msg":"async 3"`)
	test.Contains(t, sdata, `"lvl":"DBG","msg":"async 4"`)
	test.Contains(t, sdata, `"lvl":"INF","msg":"sync 1"`)
	test.Contains(t, sdata, `"lvl":"WRN","msg":"sync 2"`)
	test.Contains(t, sdata, `"lvl":"ERR","msg":"sync 3"`)
	test.Contains(t, sdata, `"msg":"context1","ctx":{"ip":"0.0.0.0"}`)
	test.Contains(t, sdata, `"msg":"context2","ctx":{"nil":null}`)
	test.Contains(t, sdata, `"msg":"context3","ctx":{"func":"unsupported field value: (func())`)
	test.Contains(t, sdata, `"msg":"context4","ctx":{"ip":"0.0.0.0"}`)
	test.Contains(t, sdata, `"msg":"context5","ctx":{"nil":null}`)
	test.Contains(t, sdata, `"msg":"context6","ctx":{"func":"unsupported field value: (func())`)
	test.Contains(t, sdata, `"msg":"context7","ctx":{"err":null}`)
	test.Contains(t, sdata, `"msg":"context8","ctx":{"err":"er1"}`)
}

func TestUnit_NewString(t *testing.T) {
	l := New()

	test.NotNil(t, l)
	l.SetFormatter(NewFormatString())

	filename, err := os.CreateTemp(os.TempDir(), "test_new_default-*.log")
	test.NoError(t, err)

	l.SetOutput(filename)
	l.SetLevel(LevelDebug)
	test.Equal(t, LevelDebug, l.GetLevel())

	go l.Infof("async %d", 1)
	go l.Warnf("async %d", 2)
	go l.Errorf("async %d", 3)
	go l.Debugf("async %d", 4)

	l.Infof("sync %d", 1)
	l.Warnf("sync %d", 2)
	l.Errorf("sync %d", 3)
	l.Debugf("sync %d", 4)

	l.WithFields(Fields{"ip": "0.0.0.0"}).Infof("context1")
	l.WithFields(Fields{"nil": nil}).Infof("context2")
	l.WithFields(Fields{"func": func() {}}).Infof("context3")

	l.WithField("ip", "0.0.0.0").Infof("context4")
	l.WithField("nil", nil).Infof("context5")
	l.WithField("func", func() {}).Infof("context6")

	l.WithError("err", nil).Infof("context7")
	l.WithError("err", fmt.Errorf("er1")).Infof("context8")

	<-time.After(time.Second * 1)
	l.Close()

	test.NoError(t, filename.Close())
	data, err := os.ReadFile(filename.Name())
	test.NoError(t, err)
	test.NoError(t, os.Remove(filename.Name()))

	sdata := string(data)
	test.Contains(t, sdata, "lvl: INF\tmsg: async 1")
	test.Contains(t, sdata, "lvl: WRN\tmsg: async 2")
	test.Contains(t, sdata, "lvl: ERR\tmsg: async 3")
	test.Contains(t, sdata, "lvl: DBG\tmsg: async 4")
	test.Contains(t, sdata, "lvl: INF\tmsg: sync 1")
	test.Contains(t, sdata, "lvl: WRN\tmsg: sync 2")
	test.Contains(t, sdata, "lvl: ERR\tmsg: sync 3")
	test.Contains(t, sdata, "lvl: DBG\tmsg: sync 4")
	test.Contains(t, sdata, "msg: context1\tctx: [[ip: 0.0.0.0]]")
	test.Contains(t, sdata, "msg: context2\tctx: [[nil: <nil>]]")
	test.Contains(t, sdata, "msg: context3\tctx: [[func: unsupported field value: (func())")
	test.Contains(t, sdata, "msg: context4\tctx: [[ip: 0.0.0.0]]")
	test.Contains(t, sdata, "msg: context5\tctx: [[nil: <nil>]]")
	test.Contains(t, sdata, "msg: context6\tctx: [[func: unsupported field value: (func())")
	test.Contains(t, sdata, "msg: context7\tctx: [[err: <nil>]]")
	test.Contains(t, sdata, "msg: context8\tctx: [[err: er1]]")
}

func BenchmarkNewJSON(b *testing.B) {
	b.ReportAllocs()

	ll := New()
	ll.SetOutput(io.Discard)
	ll.SetLevel(LevelDebug)
	ll.SetFormatter(NewFormatJSON())
	wg := sync.NewGroup()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			wg.Background(func() {
				ll.WithFields(Fields{"a": "b"}).Infof("hello")
				ll.WithField("a", "b").Infof("hello")
				ll.WithError("a", fmt.Errorf("b")).Infof("hello")
			})
		}
	})
	wg.Wait()
	ll.Close()
}

func BenchmarkNewString(b *testing.B) {
	b.ReportAllocs()

	ll := New()
	ll.SetOutput(io.Discard)
	ll.SetLevel(LevelDebug)
	ll.SetFormatter(NewFormatString())
	wg := sync.NewGroup()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			wg.Background(func() {
				ll.WithFields(Fields{"a": "b"}).Infof("hello")
				ll.WithField("a", "b").Infof("hello")
				ll.WithError("a", fmt.Errorf("b")).Infof("hello")
			})
		}
	})
	wg.Wait()
	ll.Close()
}
