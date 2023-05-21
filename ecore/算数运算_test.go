package ecore

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

func TestE四舍五入(t *testing.T) {
	gtest.AssertEQ(E四舍五入(123.144, 2), 123.14)
	gtest.AssertEQ(E四舍五入(123.145, 2), 123.15)
}

func TestE取绝对值(t *testing.T) {
	gtest.AssertEQ(E取绝对值(-100), E到数值(100))
	gtest.AssertEQ(E取绝对值(-0.1), 0.1)
	gtest.AssertEQ(E取绝对值(0.1), 0.1)
}
func TestE取整(t *testing.T) {
	gtest.AssertEQ(E取整(100.111), 100)
}

func TestE求次方(t *testing.T) {
	type args struct {
		欲求次方数值 float64
		次方数    float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			"2^2",
			args{
				2,
				2,
			},
			4,
		},
		{
			"4^2",
			args{
				4,
				2,
			},
			16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := E求次方(tt.args.欲求次方数值, tt.args.次方数); got != tt.want {
				t.Errorf("E求次方() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestE求余弦(t *testing.T) {
	type args struct {
		欲进行计算的角 float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := E求余弦(tt.args.欲进行计算的角); got != tt.want {
				t.Errorf("E求余弦() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestE求反正切(t *testing.T) {
	type args struct {
		欲求其反正切值的数值 float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := E求反正切(tt.args.欲求其反正切值的数值); got != tt.want {
				t.Errorf("E求反正切() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestE求平方根(t *testing.T) {
	type args struct {
		欲求次方数值 float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := E求平方根(tt.args.欲求次方数值); got != tt.want {
				t.Errorf("E求平方根() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestE求正切(t *testing.T) {
	type args struct {
		欲进行计算的角 float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := E求正切(tt.args.欲进行计算的角); got != tt.want {
				t.Errorf("E求正切() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestE求正弦(t *testing.T) {
	type args struct {
		欲进行计算的角 float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := E求正弦(tt.args.欲进行计算的角); got != tt.want {
				t.Errorf("E求正弦() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestE置随机数种子(t *testing.T) {
	type args struct {
		欲置入的种子数值 int64
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

// ExampleE四舍五入 用于演示E四舍五入函数的使用方法
func ExampleE四舍五入() {
	// 对3.1415926进行四舍五入，保留小数点后两位
	result := E四舍五入(3.1415926, 2)
	fmt.Println(result)
	// Output: 3.14
}

// ExampleE取绝对值 用于演示E取绝对值函数的使用方法
func ExampleE取绝对值() {
	// 求解-10.1的绝对值
	result := E取绝对值(-10.1)
	fmt.Println(result)
	// Output: 10.1
}

// ExampleE取整 用于演示E取整函数的使用方法
func ExampleE取整() {
	// 将字符串转换为int
	result := E取整("100")
	fmt.Println(result)
	// Output: 100
}

// ExampleE求次方 用于演示E求次方函数的使用方法
func ExampleE求次方() {
	// 对2的3次方进行计算
	result := E求次方(2, 3)
	fmt.Println(result)
	// Output: 8
}

// ExampleE求平方根 用于演示E求平方根函数的使用方法
func ExampleE求平方根() {
	// 求解9的平方根
	result := E求平方根(9)
	fmt.Println(result)
	// Output: 3
}

// ExampleE求正弦 用于演示E求正弦函数的使用方法
func ExampleE求正弦() {
	// 求解30度的正弦值
	result := E求正弦(30 / 180.0 * math.Pi)
	fmt.Println(result)
	// Output: 0.5
}

// ExampleE求余弦 用于演示E求余弦函数的使用方法
func ExampleE求余弦() {
	// 求解60度的余弦值
	result := E求余弦(60 / 180.0 * math.Pi)
	fmt.Println(result)
	// Output: 0.49999999999999994
}

// ExampleE求正切 用于演示E求正切函数的使用方法
func ExampleE求正切() {
	// 求解45度的正切值
	result := E求正切(45 / 180.0 * math.Pi)
	fmt.Println(result)
	// Output: 1
}

// ExampleE求反正切 用于演示E求反正切函数的使用方法
func ExampleE求反正切() {
	// 求解1的反正切值
	result := E求反正切(1)
	fmt.Println(result)
	// Output: 0.7853981633974483
}

// ExampleE置随机数种子 用于演示E置随机数种子函数的使用方法
func ExampleE置随机数种子() {
	// 使用当前时间戳作为种子数值
	E置随机数种子(0)
	// 生成10个随机数并打印
	for i := 0; i < 10; i++ {
		result := E取随机数(0, 100)
		fmt.Println(result)
	}
}

// ExampleE取随机数 用于演示E取随机数函数的使用方法
func ExampleE取随机数() {
	// 固定随机数种子
	rand.Seed(666)
	// 生成10个随机数并打印
	for i := 0; i < 10; i++ {
		result := E取随机数(0, 100)
		fmt.Println(result)
	}
}

func TestE计算数学表达式(t *testing.T) {
	t.Log("E计算数学表达式", E计算数学表达式("1+1"))
	t.Log("E计算数学表达式", E计算数学表达式("2*2"))
	t.Log("E计算数学表达式", E计算数学表达式("2/2"))
	t.Log("E计算数学表达式", E计算数学表达式("2/0"))
	t.Log("E计算数学表达式", E计算数学表达式("2>0"))
	t.Log("E计算数学表达式", E计算数学表达式("2<0"))
	t.Log("E计算数学表达式", E计算数学表达式("2-20"))
	t.Log("E计算数学表达式", E计算数学表达式("'2014-01-02' > '2014-01-01 23:59:59'"))
	t.Log("E计算数学表达式", E计算数学表达式("'2014-01-02' < '2014-01-01 23:59:59'"))
	t.Log("E计算数学表达式", E计算数学表达式("'2014-01-02' == '2014-01-02'"))

}
