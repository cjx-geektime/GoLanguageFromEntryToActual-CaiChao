package my_map

import "testing"

func TestMapWithFunValue(t *testing.T) {
	// Map 的 value 可以是一个方法
	// 这种用法可以用于创建工厂模式，key：工厂方法
	m := map[int]func(op int) int{}

	m[1] = func(op int) int {
		return op
	}

	m[2] = func(op int) int {
		return op * op
	}

	m[3] = func(op int) int {
		return op * op * op
	}

	t.Log(m[1](2), m[2](2), m[3](2))
}
