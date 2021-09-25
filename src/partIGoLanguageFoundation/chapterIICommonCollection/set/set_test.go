package set

import "testing"

// 在 Go 语言中, Set 数据结构通过 Map 实现
func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 1
	if mySet[n] {
		t.Logf("%d is exising", n)
	} else {
		t.Logf("%d is not exising", n)
	}
}
