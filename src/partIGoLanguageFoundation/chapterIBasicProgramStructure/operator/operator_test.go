package operator

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 4, 3}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}

	t.Log(a == b) // false
	//t.Log(a == c)// invalid operation: a == c (mismatched types [4]int and [5]int)
	t.Log(a == d) // true
}

func TestBitClear(t *testing.T) {
	a := 6
	t.Log(a&^1, a&^3, a&^5)// 6 4 2
}
