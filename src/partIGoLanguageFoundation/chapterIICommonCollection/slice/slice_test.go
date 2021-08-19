package slice

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	// make([]type, len, cap)
	// 其中 len 个元素会被初始化为默认值零，未初始化的元素不可以访问
	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3])
	t.Log(len(s2), cap(s2))
}

// 切片容量自增
// 容量增长规则：
/*
* 首先判断，如果新申请容量（cap）大于2倍的旧容量（old.cap），最终容量（newcap）就是新申请的容量（cap）
* 否则判断，如果旧切片的长度小于1024，则最终容量(newcap)就是旧容量(old.cap)的两倍，即（newcap=doublecap）
* 否则判断，如果旧切片长度大于等于1024，则最终容量（newcap）从旧容量（old.cap）开始循环增加原来的 1/4，即（newcap=old.cap,for {newcap += newcap/4}）直到最终容量（newcap）大于等于新申请的容量(cap)，即（newcap >= cap）
* 如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）
 */
func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i) // 如果扩容了，就会重新申请存储空间，地址会改变
		t.Log(len(s), cap(s))
	}
}

// 共享存储
func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2)
	t.Log(Q2, len(Q2), cap(Q2))

	summer := year[5:8]
	t.Log(summer)
	t.Log(summer, len(summer), cap(summer))

	summer[0] = "Jun"
	t.Log(year)
	t.Log(Q2)
	t.Log(summer)
}

// 切片的比较
func TestSliceComparing(t *testing.T) {
	var a []int
	var b = make([]int, 0, 0)
	c := []int{}
	t.Log(a, len(a), cap(a))
	t.Log(b, len(b), cap(b))
	t.Log(c, len(c), cap(c))
	t.Log(a == nil, b == nil, c == nil)// true false false
}
