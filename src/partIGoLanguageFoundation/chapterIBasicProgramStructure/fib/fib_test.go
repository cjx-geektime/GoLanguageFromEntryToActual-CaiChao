package fib

import (
	"testing"
)

func TestFibList(t *testing.T) {
	// var a int = 1
	// var b int = 1
	var (
		a int = 1
		b     = 1 //自动类型推断
	)
	// fmt.Print(a)
	t.Log(a)
	for i := 0; i < 5; i++ {
		// fmt.Print(" ", b)
		t.Log(b)
		// tmp := a
		// a = b
		// b = tmp + a
		a, b = b, a+b// 对多个变量进行同时赋值
	}
	// fmt.Println()
}
