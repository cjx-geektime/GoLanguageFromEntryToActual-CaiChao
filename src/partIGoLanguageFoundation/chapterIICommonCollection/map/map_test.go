package my_map

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d", len(m2))

	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1]) // 0, 如果取值不存在，默认返回 0，不能通过返回nil来判断元素是否存在
	m1[2] = 0
	t.Log(m1[2])
	m1[3] = 0
  // 通过下面的方法判断 map 中是否包含要取的 key
	if v, ok := m1[3]; ok {
		t.Logf("Key 3's value is %d", v)
	} else {
		t.Log("Key 3 is not extsting.")
	}
}

func TestTravelMap(t *testing.T){
  m1 := map[int]int{1: 1, 2: 4, 3: 9}
  for k, v := range m1{
    t.Log(k ,v)
  }
}
