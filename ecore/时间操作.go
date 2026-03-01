package ecore

import (
    "fmt"
    "github.com/gogf/gf/v2/os/gtime"
    "time"
)

//A simple extension for Time based on PHP's Carbon library. https://github.com/uniplaces/carbon

//到时间
//增减时间
//取时间间隔
//取某月天数
//时间到文本
//取时间部分
//取年份-
//取月份-
//取日-
//取星期几
//取小时
//取分钟
//取秒
//指定时间
//取现行时间
//置现行时间
//取日期
//取时间

// E日期时间型 日期时间类型，提供日期时间的相关操作
type E日期时间型 struct {
    Time *gtime.Time
}

// New日期时间型 创建一个新的日期时间对象，初始值为当前时间
//
// 返回值：
//   *E日期时间型 - 日期时间对象指针
func New日期时间型() *E日期时间型 {
    return &E日期时间型{Time: gtime.Now()}
}

// E取星期几 获取当前日期是星期几（1-7，1表示星期一）
//
// 返回值：
//   int64 - 星期几
func (this *E日期时间型) E取星期几() int64 {
    return E到整数(this.Time.Format("N"))
}

// E取月天数 获取当前月份的天数
//
// 返回值：
//   int64 - 月份天数
func (this *E日期时间型) E取月天数() int64 {
    return E到整数(this.Time.Format("t"))
}

// E取年份 获取年份
//
// 返回值：
//   int64 - 年份
func (this *E日期时间型) E取年份() int64 {
    return E到整数(this.Time.Format("Y"))
}

// E取月份 获取月份（1-12）
//
// 返回值：
//   int64 - 月份
func (this *E日期时间型) E取月份() int64 {
    return E到整数(this.Time.Format("m"))
}

// E取日 获取日期（1-31）
//
// 返回值：
//   int64 - 日期
func (this *E日期时间型) E取日() int64 {
    return E到整数(this.Time.Format("d"))
}

// E取小时 获取小时（0-23）
//
// 返回值：
//   int64 - 小时
func (this *E日期时间型) E取小时() int64 {
    return E到整数(this.Time.Format("H"))
}

// E取分钟 获取分钟（0-59）
//
// 返回值：
//   int64 - 分钟
func (this *E日期时间型) E取分钟() int64 {
    return E到整数(this.Time.Format("i"))
}

// E取秒 获取秒（0-59）
//
// 返回值：
//   int64 - 秒
func (this *E日期时间型) E取秒() int64 {
    return E到整数(this.Time.Format("s"))
}

// E取毫秒 获取毫秒
//
// 返回值：
//   int64 - 毫秒
func (this *E日期时间型) E取毫秒() int64 {
    return E到整数(this.Time.Millisecond())
}

// E取微秒 获取微秒
//
// 返回值：
//   int64 - 微秒
func (this *E日期时间型) E取微秒() int64 {
    return E到整数(this.Time.Microsecond())
}

// E取纳秒 获取纳秒
//
// 返回值：
//   int64 - 纳秒
func (this *E日期时间型) E取纳秒() int64 {
    return E到整数(this.Time.Nanosecond())
}

// E取时间戳 获取Unix时间戳（秒）
//
// 返回值：
//   int64 - Unix时间戳
func (this *E日期时间型) E取时间戳() int64 {
    return this.Time.Timestamp()
}

// E取时间戳毫秒 获取Unix时间戳（毫秒）
//
// 返回值：
//   int64 - Unix时间戳（毫秒）
func (this *E日期时间型) E取时间戳毫秒() int64 {
    return this.Time.TimestampMilli()
}

// E取时间戳微秒 获取Unix时间戳（微秒）
//
// 返回值：
//   int64 - Unix时间戳（微秒）
func (this *E日期时间型) E取时间戳微秒() int64 {
    return this.Time.TimestampMicro()
}

// E取时间戳纳秒 获取Unix时间戳（纳秒）
//
// 返回值：
//   int64 - Unix时间戳（纳秒）
func (this *E日期时间型) E取时间戳纳秒() int64 {
    return this.Time.TimestampNano()
}

// E时间到文本 将日期时间转换为文本格式
//
// 参数：
//   format string - 格式化字符串，默认为"Y-m-d H:i:s"
//
// 返回值：
//   string - 格式化后的时间文本
func (this *E日期时间型) E时间到文本(format string) string {
    if format == "" {
        format = "Y-m-d H:i:s"
    }
    return this.Time.Format(format)
}

// E增减日期 增加或减少日期
//
// 参数：
//   年 int - 要增加的年数（可为负数）
//   月 int - 要增加的月数（可为负数）
//   日 int - 要增加的天数（可为负数）
//
// 返回值：
//   *E日期时间型 - 日期时间对象本身，支持链式调用
func (this *E日期时间型) E增减日期(年 int, 月 int, 日 int) *E日期时间型 {
    this.Time = this.Time.AddDate(年, 月, 日)
    return this
}

// E增减时间 增加或减少时间
//
// 参数：
//   时 int - 要增加的小时数（可为负数）
//   分 int - 要增加的分钟数（可为负数）
//   秒 int - 要增加的秒数（可为负数）
//
// 返回值：
//   *E日期时间型 - 日期时间对象本身，支持链式调用
func (this *E日期时间型) E增减时间(时 int, 分 int, 秒 int) *E日期时间型 {
    if 时 != 0 {
        this.Time.AddStr(E到文本(时) + "h")
    }
    if 分 != 0 {
        this.Time.AddStr(E到文本(分) + "m")
    }
    if 秒 != 0 {
        this.Time.AddStr(E到文本(秒) + "s")
    }
    return this
}

// E大于 判断当前时间是否在指定时间之后
//
// 参数：
//   time *E日期时间型 - 要比较的日期时间
//
// 返回值：
//   bool - 是返回true，否则返回false
func (this *E日期时间型) E大于(time *E日期时间型) bool {
    return this.Time.After(time.Time)
}

// E小于 判断当前时间是否在指定时间之前
//
// 参数：
//   time *E日期时间型 - 要比较的日期时间
//
// 返回值：
//   bool - 是返回true，否则返回false
func (this *E日期时间型) E小于(time *E日期时间型) bool {
    return this.Time.Before(time.Time)
}

// E等于 判断当前时间是否与指定时间相等
//
// 参数：
//   time *E日期时间型 - 要比较的日期时间
//
// 返回值：
//   bool - 相等返回true，否则返回false
func (this *E日期时间型) E等于(time *E日期时间型) bool {
    return this.Time.Equal(time.Time)
}

// E到友好时间
// 返回当前时间和调用对象时间的友好时间差异。
// 如果调用对象的时间晚于当前时间，则返回时间差异的可读格式：
// n分钟后，n个月后，或者很久以后。
// 如果调用对象的时间早于当前时间，则返回时间差异的可读格式：
// n分钟前，n个月前，或者很久以前。
// @return  时间差异的友好可读格式
func (this *E日期时间型) E到友好时间() string {
    t := this.Time.Time
    now := time.Now()
    duration := now.Sub(t)
    if duration < 0 {
        duration = -duration
        if duration < time.Minute {
            return fmt.Sprintf("%d 秒后", int(duration.Seconds()))
        } else if duration < time.Hour {
            return fmt.Sprintf("%d 分钟后", int(duration.Minutes()))
        } else if duration < time.Hour*24 {
            return fmt.Sprintf("%d 小时后", int(duration.Hours()))
        } else if duration < time.Hour*24*30 {
            return fmt.Sprintf("%d 天后", int(duration.Hours()/24))
        } else if duration < time.Hour*24*365 {
            return fmt.Sprintf("%d 个月后", int(duration.Hours()/24/30))
        } else {
            return "很久以后"
        }
    } else {
        if duration < time.Minute {
            return fmt.Sprintf("%d 秒前", int(duration.Seconds()))
        } else if duration < time.Hour {
            return fmt.Sprintf("%d 分钟前", int(duration.Minutes()))
        } else if duration < time.Hour*24 {
            return fmt.Sprintf("%d 小时前", int(duration.Hours()))
        } else if duration < time.Hour*24*30 {
            return fmt.Sprintf("%d 天前", int(duration.Hours()/24))
        } else if duration < time.Hour*24*365 {
            return fmt.Sprintf("%d 个月前", int(duration.Hours()/24/30))
        } else {
            return "很久以前"
        }
    }
}

// E取现行时间 获取当前系统时间
//
// 返回值：
//   *E日期时间型 - 当前时间的日期时间对象
func E取现行时间() *E日期时间型 {
    this := new(E日期时间型)
    this.Time = gtime.Now()
    return this
}

// E到时间 将文本转换为日期时间对象
//
// 参数：
//   s string - 时间文本，如"2023-01-01 12:00:00"
//
// 返回值：
//   *E日期时间型 - 日期时间对象
func E到时间(s string) *E日期时间型 {
    this := new(E日期时间型)
    if t, err := gtime.StrToTime(s); err == nil {
        this.Time = t
    }
    return this
}

// E到时间从时间戳 将Unix时间戳转换为日期时间对象
//
// 参数：
//   s int64 - Unix时间戳
//
// 返回值：
//   *E日期时间型 - 日期时间对象
func E到时间从时间戳(s int64) *E日期时间型 {
    this := new(E日期时间型)
    this.Time = gtime.NewFromTimeStamp(s)
    return this
}

// E取时间戳 获取当前系统的时间戳（秒）
//
// 返回值：
//   int64 - Unix时间戳
func E取时间戳() int64 {
    return gtime.Now().Timestamp()
}
