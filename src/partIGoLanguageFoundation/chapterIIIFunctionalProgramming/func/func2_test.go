package fn_test

import (
	"fmt"
	"testing"
	"time"
)

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 2)
	return op
}

// 计算函数的执行时间
func TestFn2(t *testing.T) {
	tsSf := timeSpent(slowFun)
	t.Log(tsSf(10))
}
