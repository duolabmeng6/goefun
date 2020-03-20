package E

import (
	"testing"
)

func Test_取现行时间(t *testing.T) {
	time := E取现行时间()
	t.Log("取年份", time.E取年份())
	t.Log("E取月份", time.E取月份())
	t.Log("E取日", time.E取日())
	t.Log("E取星期几", time.E取星期几())
	t.Log("E取小时", time.E取小时())
	t.Log("E取分钟", time.E取分钟())
	t.Log("E取秒", time.E取秒())
	t.Log("E取月天数", time.E取月天数())
	t.Log("现行时间", time.E时间到文本("Y-m-d H:i:s"))
	time = E到时间("2019-02-04 22:53:02")
	t.Log("取年份", time.E取年份())
	t.Log("E取月份", time.E取月份())
	t.Log("E取日", time.E取日())
	t.Log("E取星期几", time.E取星期几())
	t.Log("E取小时", time.E取小时())
	t.Log("E取分钟", time.E取分钟())
	t.Log("E取秒", time.E取秒())
	t.Log("E取月天数", time.E取月天数())
	t.Log("现行时间", time.E时间到文本("Y-m-d H:i:s"))

	time.E增减日期(1, 0, 0)
	t.Log("E增减时间", time.E时间到文本("Y-m-d H:i:s"))

	time.E增减时间(1, 1, 1)
	t.Log("E增减时间", time.E时间到文本("Y-m-d H:i:s"))

	time.E增减时间(-1, -1, -1)
	t.Log("E增减时间", time.E时间到文本("Y-m-d H:i:s"))

	time2 := E到时间("2022-03-04 22:53:02")
	t.Log("E大于", time.E大于(time2))
	t.Log("E小于", time.E小于(time2))

	time3 := E到时间("E增减时间 2020-03-04 22:53:02")
	t.Log("E等于", time.E等于(time3))

	time4 := E取现行时间()
	t.Log("E取时间戳", time4.E取时间戳())
	t.Log("E取时间戳毫秒", time4.E取时间戳毫秒())
	t.Log("E取时间戳微秒", time4.E取时间戳微秒())
	t.Log("E取时间戳纳秒", time4.E取时间戳纳秒())

	t.Log("E到友好时间", time4.E到友好时间(E取现行时间()))

	t.Log("E到时间从时间戳", E到时间从时间戳(E取现行时间().E取时间戳()).E时间到文本(""))

	t.Log("E取时间戳", E取现行时间().E取时间戳())

}
