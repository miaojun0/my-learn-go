package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Println("pet")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(host)
}

type Dog struct {
	Pet
}

func (d *Dog) Speak() {
	fmt.Println("dog")
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("dog")
}
