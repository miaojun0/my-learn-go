package _interface

import (
	"fmt"
	"testing"
)

type Programmer interface {
	WriteCode() string
}

type GoProgrammer struct{}

func (*GoProgrammer) WriteCode() string {
	return fmt.Sprintf("hello world")
}

func TestClient(t *testing.T) {
	p := new(GoProgrammer)
	ret := p.WriteCode()
	fmt.Println(ret)
}
