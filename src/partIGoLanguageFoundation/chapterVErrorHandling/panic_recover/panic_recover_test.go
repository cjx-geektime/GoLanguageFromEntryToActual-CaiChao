package panic_recover

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	// 可以配合 panic() 当做 finally 使用
	defer func() {
		//  recover(): 错误恢复
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
	}()
	fmt.Println("Start")
	// 抛出错误，最终会执行 defer 方法（输出当前调用用栈信息）
	panic(errors.New("Something wrong!"))
	// 退出程序，不会执行 defer 方法（不输出当前调用用栈信息）
	// os.Exit(-1)
	// fmt.Println("End")
}
