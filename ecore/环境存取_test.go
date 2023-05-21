package ecore

import (
	"fmt"
	"os"
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
	E写环境变量("test", "123")
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

func ExampleE取命令行() {
	os.Args = []string{"example", "foo", "bar", "baz"}
	args := E取命令行()
	fmt.Println(args)
	// Output: [example foo bar baz]
}
func ExampleE取运行目录() {
	path := E取运行目录()
	fmt.Println(path)
}

func ExampleE读环境变量() {
	os.Setenv("MY_ENV_VAR", "my_value")
	value := E读环境变量("MY_ENV_VAR")
	fmt.Println(value)
	//  Output:  my_value
}

func ExampleE写环境变量() {
	result := E写环境变量("MY_ENV_VAR", "my_new_value")
	fmt.Println(result)
	//  Output:  true
}

func TestE设置命令行(t *testing.T) {
	E调试输出(E取命令行()[1])
	E调试输出(E取命令行()[2])
	E调试输出(E取命令行()[3])

	var a, b, c string
	E设置命令行("a", "a", "这是一个a参数", &a)
	E设置命令行("b", "b", "这是一个b参数", &b)
	E设置命令行("c", "c", "这是一个c参数", &c)
	E命令行解析()

	E调试输出(a, b, c)
}

func TestE加载环境变量文件(t *testing.T) {
	E加载环境变量_从文件("test.env")
	t.Log("SECRET_KEY", E读环境变量("S3_BUCKET"))
	t.Log("SECRET_KEY", E读环境变量("SECRET_KEY"))
	t.Log("default", E读环境变量("default", ""))
	t.Log("default", E读环境变量("default", "123"))

	t.Log("default", Env("default", "123"))

	env, _ := E加载环境变量_从内容到map(E到文本(E读入文件("test.env")))
	for k, v := range env {
		t.Log("env", k, v)
	}

}
