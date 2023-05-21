package ecore

import (
	"math"
	"math/rand"
	"time"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"

	"github.com/Knetic/govaluate"
)

// E四舍五入 用于对欲被四舍五入的数值进行四舍五入操作
//
// 参数：
//
//	欲被四舍五入的数值：float64，需要执行四舍五入的数值
//	被舍入的位置：int，需要执行四舍五入的小数位置
//
// 返回值：
//
//	float64：四舍五入后的结果
func E四舍五入(欲被四舍五入的数值 float64, 被舍入的位置 int) float64 {
	var pow float64 = 1
	for i := 0; i < 被舍入的位置; i++ {
		pow *= 10
	}
	return float64(int((欲被四舍五入的数值*pow)+0.5)) / pow
}

// E取绝对值 用于对双精度小数型进行取绝对值操作
//
// 参数：
//
//	value：float64，需要执行取绝对值操作的数值
//
// 返回值：
//
//	float64：取绝对值后的结果
func E取绝对值(双精度小数型 float64) float64 {
	return math.Abs(双精度小数型)
}

// E取整 用于对各种类型进行执行取整操作
//
// 参数：
//
//	value：interface{}，需要执行取整操作的数值（支持int、int64、float32、float64和string类型）
//
// 返回值：
//
//	int：取整后的结果
func E取整(value interface{}) int {
	return gconv.Int(value)
}

// E求次方 用于求欲求次方数值的次方
//
// 参数：
//
//	欲求次方数值：float64，需要执行求次方操作的数值
//	次方数：float64，需要进行的次方数
//
// 返回值：
//
//	float64：求次方的结果
func E求次方(欲求次方数值 float64, 次方数 float64) float64 {
	return math.Pow(欲求次方数值, 次方数)
}

// E求平方根 用于求欲求平方根的数值
//
// 参数：
//
//	欲求平方根的数值：float64，需要执行求平方根操作的数值
//
// 返回值：
//
//	float64：求平方根的结果
func E求平方根(欲求平方根的数值 float64) float64 {
	return math.Sqrt(欲求平方根的数值)
}

// E求正弦 用于对欲进行计算的角的正弦值求解
//
// 参数：
//
//	欲进行计算的角：float64，需要求解正弦值的角度值
//
// 返回值：
//
//	float64：角的正弦值
func E求正弦(欲进行计算的角 float64) float64 {
	return math.Sin(欲进行计算的角)
}

// E求余弦 用于对欲进行计算的角的余弦值求解
//
// 参数：
//
//	欲进行计算的角：float64，需要求解余弦值的角度值
//
// 返回值：
//
//	float64：角的余弦值
func E求余弦(欲进行计算的角 float64) float64 {
	return math.Cos(欲进行计算的角)
}

// E求正切 用于对欲进行计算的角的正切值求解
//
// 参数：
//
//	欲进行计算的角：float64，需要求解正切值的角度值
//
// 返回值：
//
//	float64：角的正切值
func E求正切(欲进行计算的角 float64) float64 {
	return math.Tan(欲进行计算的角)
}

// E求反正切 用于对欲求其反正切值的数值求解
//
// 参数：
//
//	欲求其反正切值的数值：float64，需要求解反正切值的数值
//
// 返回值：
//
//	float64：数值的反正切值
func E求反正切(欲求其反正切值的数值 float64) float64 {
	return math.Atan(欲求其反正切值的数值)
}

// E置随机数种子 用于设置随机数种子
//
// 参数：
//
//	欲置入的种子数值：int64，需要设置的种子数值，如果为0则使用当前时间戳作为种子数值
func E置随机数种子(欲置入的种子数值 int64) {
	if 欲置入的种子数值 == 0 {
		欲置入的种子数值 = time.Now().UnixNano()
	}
	rand.Seed(欲置入的种子数值)
}

// E取随机数 用于获取指定范围内的随机整数
//
// 参数：
//
//	欲取随机数的最小值：int，需要获取的随机数的最小值（包含在内）
//	欲取随机数的最大值：int，需要获取的随机数的最大值（不包含在内）
//
// 返回值：
//
//	int：随机数值
func E取随机数(欲取随机数的最小值 int, 欲取随机数的最大值 int) int {
	return grand.N(欲取随机数的最小值, 欲取随机数的最大值)
}

//更多的表达式用法请看这个
//https://github.com/Knetic/govaluate

//expression, err := govaluate.NewEvaluableExpression("(requests_made * requests_succeeded / 100) >= 90");
//
//parameters := make(map[string]interface{}, 8)
//parameters["requests_made"] = 100;
//parameters["requests_succeeded"] = 80;
//
//result, err := expression.Evaluate(parameters);

// 子程序名：计算数学表达式
// 返回值类型：双精度小数型
// 参数<1>的名称为“数学表达式”，类型为“文本型”。
func E计算数学表达式(v string) string {
	expression, err := govaluate.NewEvaluableExpression(v)
	if err != nil {
		return ""
	}
	result, err := expression.Eval(nil)
	if err != nil {
		return ""
	}
	// result is now set to "50.0", the float64 value.
	return E到文本(result)
}
