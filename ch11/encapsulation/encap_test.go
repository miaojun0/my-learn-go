package encapsulation

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   int
	Name string
	Age  int
}

//func (e *Employee) String() string {
//	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
//	return fmt.Sprintf("ID: %d-Name: %s-Age: %d", e.Id, e.Name, e.Age)
//}

func (e Employee) String() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID: %d-Name: %s-Age: %d", e.Id, e.Name, e.Age)
}

func TestCreateEmployee(t *testing.T) {
	e1 := Employee{1, "jhon", 23}
	e2 := Employee{Id: 2, Name: "mike", Age: 26}
	t.Log(e1, e2)

	var e3 Employee
	t.Log(e3)

	e4 := new(Employee)
	t.Log(e4)

	var e5 *Employee = nil
	t.Log(e5) //nil
}

func TestStructOperate(t *testing.T) {
	e := Employee{1, "jhon", 23}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
}
