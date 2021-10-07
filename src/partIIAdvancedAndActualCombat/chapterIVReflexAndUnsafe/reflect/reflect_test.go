package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTypeAndValue(t *testing.T) {
	var f int64 = 10
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
}

//检查反射类型
//用空接口接收任意类型
func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("Integer")
	default:
		fmt.Println("Unknow", t)
	}
}

func TestCheckType(t *testing.T) {
	var f float64 = 32
	CheckType(f)
	CheckType(&f)
}

//通过反射调用结构体
type Employee struct {
	EmployeeID string
	Name       string `format:"normal"` // struct tag
	Age        int
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{"1", "Mike", 30}
	// 按名字访问结构的成员
	t.Logf("Name: value(%[1]v), Type(%[1]T) ", reflect.ValueOf(*e).FieldByName("Name"))
	// 访问 StructTag
	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("Failed to get 'name' field.")
	} else {
		t.Log("Tag:format", nameField.Tag.Get("format"))
	}
	// 按名字访问结构的方方法
	reflect.ValueOf(e).MethodByName("UpdateAge").
		Call([]reflect.Value{reflect.ValueOf(31)})
	t.Log("Updated Age: ", e)
}
