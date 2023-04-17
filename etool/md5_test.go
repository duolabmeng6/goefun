package E

import (
	"testing"

	E "github.com/duolabmeng6/goefun/ecore"
)

func TestE取md5从文本(t *testing.T) {
	t.Log("E取md5从文本", E取md5从文本("123456"))
	t.Log("E取md5", E取md5(E.E到字节集("123456")))

}
