package set

import "testing"

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
