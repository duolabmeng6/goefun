package E

import "testing"

func TestE删全部空(t *testing.T) {
	t.Log("E删全部空", E删全部空("我 爱 e f u n 好   棒"))
}
func TestE子文本替换(t *testing.T) {

	t.Log("TestE子文本替换", E子文本替换("作    者：颜温","作    者：",""))
}

