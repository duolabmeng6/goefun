package E

import (
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
