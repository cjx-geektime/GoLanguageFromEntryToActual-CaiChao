package empty_interface

import (
	"fmt"
	"testing"
)

type StructType struct{}

// 空接口可以接收所有类型的数据
func Assert(p interface{}) {
  // 第二个返回值是操作结果：true or false
	if i, ok := p.(int); ok {
    println(ok)
		fmt.Println("Integer", i)
		return
	}else{
    println(ok)
  }

	if s, ok := p.(string); ok {
    println(ok)
		fmt.Println("String", s)
		return
	}else{
    println(ok)
  }

	fmt.Printf("Unknow Type %T", p)
	fmt.Println()
}

func Assert2(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Printf("Unknow Type %T", v)
		fmt.Println()
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	Assert(10)
	Assert("10")
	Assert(StructType{})

	Assert2(10)
	Assert2("10")
	Assert2(StructType{})
}
