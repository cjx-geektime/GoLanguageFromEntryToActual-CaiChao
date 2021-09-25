package set

import (
	"fmt"
	"testing"
)

// 在 Go 语言中, Set 数据结构通过 Map 实现
func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	setExist(mySet, 1)
	t.Log(len(mySet))
	mySet[3] = false
	t.Log(len(mySet))
	mySet[3] = true
	t.Log(len(mySet))
	delete(mySet, 1)
	t.Log(len(mySet))
}

func setExist(mySet map[int]bool, key int) {
	if mySet[key] {
		fmt.Println("exising")
	} else {
		fmt.Println("not exising")
	}
}

func setGet(mySet map[int]bool, key int) {
	if mySet[key] {
		fmt.Println(mySet[key])
	} else {
		fmt.Println("not exising")
	}
}
