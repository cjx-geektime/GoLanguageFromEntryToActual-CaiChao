package data_type

import "testing"

// 自定义数据类型，这里仅仅是起别名
type MyInt int64

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(a)

	t.Log(a, b, c)
}
