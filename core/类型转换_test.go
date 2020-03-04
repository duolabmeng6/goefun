package core

import (
	"reflect"
	"testing"
)

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
			"test",
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
			"test",
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
			"test",
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

func TestE到整数(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"test",
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

func TestE到文本(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"test",
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

func TestE文本编码转换(t *testing.T) {
	type args struct {
		src        string
		oldEncoder string
		newEncoder string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := E文本编码转换(tt.args.src, tt.args.oldEncoder, tt.args.newEncoder); got != tt.want {
				t.Errorf("E文本编码转换() = %v, want %v", got, tt.want)
			}
		})
	}
}