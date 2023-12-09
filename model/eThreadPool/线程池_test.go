package eThreadPool

import (
	"fmt"
	"github.com/duolabmeng6/goefun/ecore"
	"testing"
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
	fmt.Println("ok")
}

var (
	myMap = make(map[int]int, 10)
	lock  E读写共享锁
)

func worker(id int) {
	fmt.Printf("worker:%d start \n", id)
	ecore.E延时(1000)
	lock.E开始写()
	myMap[id] = id * 2
	lock.E结束写()

	lock.E开始读()
	fmt.Println(myMap[id])
	lock.E结束读()
}
