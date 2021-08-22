package fn_test

import "testing"

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3))
	t.Log(Sum(2, 3, 4, 5))
}
