/*
 *  Copyright (c) 2019-2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package version

import (
	"reflect"
	"testing"
)

func TestUnit_Parse(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    *Version
		wantErr bool
	}{
		{
			name: "Case1",
			args: "v1.1000.1231",
			want: &Version{
				Major: 1,
				Minor: 1000,
				Patch: 1231,
			},
			wantErr: false,
		},
		{
			name: "Case2",
			args: "app/v1.1000.1231",
			want: &Version{
				Major: 1,
				Minor: 1000,
				Patch: 1231,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnit_Max(t *testing.T) {
	tests := []struct {
		name    string
		vers    []string
		wantOut string
	}{
		{
			name:    "Case1",
			vers:    []string{"v0.0.1", "v0.0.119991"},
			wantOut: "v0.0.119991",
		},
		{
			name:    "Case2",
			vers:    []string{},
			wantOut: "v0.0.0",
		},
		{
			name:    "Case3",
			vers:    []string{" "},
			wantOut: "v0.0.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := Max(tt.vers...); gotOut.String() != tt.wantOut {
				t.Errorf("Max() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
