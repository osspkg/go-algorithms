/*
 *  Copyright (c) 2019-2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package network_test

import (
	"reflect"
	"testing"

	"go.osspkg.com/x/network"
)

func TestUnit_Normalize(t *testing.T) {
	tests := []struct {
		name string
		port string
		args []string
		want []string
	}{
		{
			name: "Case1",
			port: "53",
			args: []string{"1.1.1.1", "1.1.1.1:123", "123.11.11"},
			want: []string{"1.1.1.1:53", "1.1.1.1:123"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := network.Normalize(tt.port, tt.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}
