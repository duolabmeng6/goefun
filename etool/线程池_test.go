package E

import (
	"fmt"
	"testing"

	. "github.com/duolabmeng6/goefun/ecore"
)

func TestNew线程池(t *testing.T) {
	pool := New线程池(20)
	for i := 0; i < 20; i++ {
		pool.E加入任务()
		go func(i int) {
			defer pool.E完成()
			worker(i)
		}(i)
	}
	pool.E等待任务完成()
	E调试输出("ok")
}

var (
	myMap = make(map[int]int, 10)
	lock  E读写共享锁
)

func worker(id int) {
	fmt.Printf("worker:%d start \n", id)
	E延时(1000)
	lock.E开始写()
	myMap[id] = id * 2
	lock.E结束写()

	lock.E开始读()
	E调试输出(myMap[id])
	lock.E结束读()
}
