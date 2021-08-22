package fn_test

import (
	"fmt"
	"testing"
)

func Clear() {
	fmt.Println("Clear resources.")
}

func TestDefer(t *testing.T) {
	defer Clear() // 函数结束前必须执行，类似Java中的finaly
	fmt.Println("Start")
	panic("Error!") // 抛出错误
}
