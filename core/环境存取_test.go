package core

import (
	"reflect"
	"testing"
)

func TestE写环境变量(t *testing.T) {
	type args struct {
		环境变量名称 string
		欲写入内容  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test",
			args{
				"test",
				"123",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := E写环境变量(tt.args.环境变量名称, tt.args.欲写入内容); got != tt.want {
				t.Errorf("E写环境变量() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestE取命令行(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := E取命令行(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("E取命令行() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestE取运行目录(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := E取运行目录(); got != tt.want {
				t.Errorf("E取运行目录() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestE读环境变量(t *testing.T) {
	E写环境变量("test","123")
	type args struct {
		环境变量名称 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test",
			args{
				"test",
			},
			"123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := E读环境变量(tt.args.环境变量名称); got != tt.want {
				t.Errorf("E读环境变量() = %v, want %v", got, tt.want)
			}
		})
	}
}