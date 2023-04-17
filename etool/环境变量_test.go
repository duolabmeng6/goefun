package E

import (
	"testing"

	. "github.com/duolabmeng6/goefun/ecore"
)

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
