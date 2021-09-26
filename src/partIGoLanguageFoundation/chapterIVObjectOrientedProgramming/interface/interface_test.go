package interface_test

import "testing"

type Programmer interface {
	WriteHelloWorld() string
	SayHello() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func (g *GoProgrammer) SayHello() string {
	return "fmt.Println(\"Hello\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
	t.Log(p.SayHello())
}
