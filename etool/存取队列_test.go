package etool

import (
	"context"
	"github.com/duolabmeng6/goefun/ecore"
	"testing"
	"time"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/os/gtimer"
)

func TestNew存取队列(t *testing.T) {
	var (
		ctx = gctx.New()
	)

	q := New存取队列()
	// 数据生产者，每隔1秒往队列写数据
	gtimer.SetTimeout(ctx, time.Second, func(ctx context.Context) {
		var v = gtime.Now().String()
		q.E压入队列(v)
		ecore.E调试输出("Push:", v)
	})

	// 3秒后关闭队列
	gtimer.SetTimeout(ctx, 5*time.Second, func(ctx context.Context) {
		q.E清空()
	})

	// 消费者，不停读取队列数据并输出到终端
	for {
		if v := q.E弹出队列(); v != nil {
			ecore.E调试输出(" Pop:", v)
		} else {
			break
		}
	}
}
