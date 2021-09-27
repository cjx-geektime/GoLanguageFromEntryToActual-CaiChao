package object_pool

import (
	"fmt"
	"testing"
	"time"
	"unsafe"
)

func TestObjPool(t *testing.T) {
	pool := NewObjPool(10)
	for i := 0; i < 11; i++ {
		if v, err := pool.getObj(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			fmt.Printf("%T %x\n", v, unsafe.Pointer(v))
			if err := pool.ReleaseObj(v); err != nil {
				t.Error(err)
			}
		}
	}
	fmt.Println("Done")
}
