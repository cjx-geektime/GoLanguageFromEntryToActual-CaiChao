package groutine_test

import (
	"fmt"
	"testing"
	"time"
)

func TestGroutine(t *testing.T) {
	for i := 0; i < 20; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	// 如果不休眠一会的话，主进程会在协程之前退出
	time.Sleep(time.Millisecond * 50)
}
