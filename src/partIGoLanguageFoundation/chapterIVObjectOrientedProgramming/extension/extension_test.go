package extension

import (
	"fmt"
	"testing"
)

type Pet struct{}

func (p *Pet) Speak() {
	fmt.Println("Pet")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	Pet
}

type Dog2 struct {
	p *Pet
}

func (d *Dog) Speak() {
	fmt.Println("Dog")
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.Speak()
	dog.SpeakTo("C")

	dog2 := new(Dog2)
	dog2.p.Speak()
}
