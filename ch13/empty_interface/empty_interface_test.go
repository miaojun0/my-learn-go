package empty_interface

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	//if i, ok := p.(int); ok {
	//	fmt.Println("int", i)
	//} else {
	//	fmt.Println(ok)
	//}

	switch v := p.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	}
}

func TestClient(t *testing.T) {
	DoSomething("a")
}
