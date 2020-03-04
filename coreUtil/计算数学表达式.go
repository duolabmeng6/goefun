package coreUtil

import (
	"github.com/Knetic/govaluate"
	. "github.com/duolabmeng6/goefun/core"
)
//更多的表达式用法请看这个
//https://github.com/Knetic/govaluate

//expression, err := govaluate.NewEvaluableExpression("(requests_made * requests_succeeded / 100) >= 90");
//
//parameters := make(map[string]interface{}, 8)
//parameters["requests_made"] = 100;
//parameters["requests_succeeded"] = 80;
//
//result, err := expression.Evaluate(parameters);

//子程序名：计算数学表达式
//返回值类型：双精度小数型
//参数<1>的名称为“数学表达式”，类型为“文本型”。
func E计算数学表达式(v string) string{
	expression, err := govaluate.NewEvaluableExpression(v);
	if err!=nil{
		return ""
	}
	result, err := expression.Eval(nil);
	if err!=nil{
		return ""
	}
	// result is now set to "50.0", the float64 value.
	return E到文本(result)
}