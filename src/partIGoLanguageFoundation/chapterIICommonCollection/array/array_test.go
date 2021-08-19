package array

import "testing"

// 定义数组
func TestArrayInit(t *testing.T) {
	var arr [3]int             // 声明并初始化，默认值为0
	arr1 := [4]int{1, 2, 3, 4} // 声明同时初始化
	arr2 := [...]int{1, 2, 3, 5}
	arr1[1] = 5
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr2)

	c := [2][2]int{{1, 2}, {3, 4}} // 多维数组初始化
	t.Log(c)
}

// 遍历数组
func TestArrayTravel(t *testing.T) {
	arr := [...]int{1, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}
	for idx, e := range arr {
		t.Log(idx, e)
	}

	//如果不想用 idx 就可以用 _ 代替，表示我们不 care idx
	for _, e := range arr {
		t.Log(e)
	}
}

// 数组截取
func TestArraySection(t *testing.T) {
	//a[开始索引(包含), 结束索引(不包含)]
	a := [...]int{1, 2, 3, 4, 5}
	t.Log(a[1:2])
	t.Log(a[1:3])
	t.Log(a[1:len(a)])
	t.Log(a[1:])
	t.Log(a[:2])
	t.Log(a[:])
}
