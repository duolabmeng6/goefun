// Copyright 2023 The duolabmeng6 Authors. All rights reserved.
// license that can be found in the LICENSE file.

package ecore

import (
	"reflect"
	"testing"
)

// TestE到字节 测试E到字节函数
func TestE到字节(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			"测试整数转换",
			args{
				65,
			},
			65,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := E到字节(tt.args.value); got != tt.want {
				t.Errorf("E到字节() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestE到字节集 测试E到字节集函数
func TestE到字节集(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"测试字符串转换",
			args{
				"A",
			},
			[]byte{65},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := E到字节集(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("E到字节集() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestE到数值 测试E到数值函数
func TestE到数值(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"测试字符串转换",
			args{
				"123.123",
			},
			123.123,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := E到数值(tt.args.value); got != tt.want {
				t.Errorf("E到数值() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestE到整数 测试E到整数函数
func TestE到整数(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"测试字符串转换",
			args{
				"123.123",
			},
			123,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := E到整数(tt.args.value); got != tt.want {
				t.Errorf("E到整数() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestE到文本 测试E到文本函数
func TestE到文本(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"测试整数转换",
			args{
				123,
			},
			"123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := E到文本(tt.args.value); got != tt.want {
				t.Errorf("E到文本() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestE到结构体 测试E到结构体函数
func TestE到结构体(t *testing.T) {
	type args struct {
		待转换的参数 interface{}
		结构体指针  interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := E到结构体(tt.args.待转换的参数, tt.args.结构体指针); (err != nil) != tt.wantErr {
				t.Errorf("E到结构体() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
