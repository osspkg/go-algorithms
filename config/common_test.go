/*
 *  Copyright (c) 2019-2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package config

import (
	"os"
	"testing"

	"go.osspkg.com/x/test"
)

type (
	testConfigItem struct {
		Home string `yaml:"home"`
		Path string `yaml:"path"`
	}
	testConfig struct {
		Envs testConfigItem `yaml:"envs"`
	}
)

func TestUnit_ConfigResolve(t *testing.T) {
	filename := "/tmp/TestUnit_ConfigResolve.yaml"
	data := `
envs:
  home: "@env(HOME#fail)"
  path: "@env(PATH#fail)"
`
	err := os.WriteFile(filename, []byte(data), 0755)
	test.NoError(t, err)

	res := New(EnvResolver())

	err = res.OpenFile(filename)
	test.NoError(t, err)
	err = res.Build()
	test.NoError(t, err)

	tc := &testConfig{}

	err = res.Decode(tc)
	test.NoError(t, err)
	test.NotEqual(t, "fail", tc.Envs.Home)
	test.NotEqual(t, "fail", tc.Envs.Path)
}
