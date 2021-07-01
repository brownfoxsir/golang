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

//func (e Employee) String() string {
//	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
//	return fmt.Sprintf("ID:%d-Name:%s-Age:%d", e.Id, e.Name, e.Age)
//}

func (e *Employee) String() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%d/Name:%s/Age:%d", e.Id, e.Name, e.Age)
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{
		Id:   1,
		Name: "Edge",
		Age:  10,
	}
	e1 := Employee{2, "Plan", 11}
	e2 := new(Employee) //注意这⾥返回的引⽤/指针，相当于 e := &Employee{}
	e2.Id = 2
	e2.Name = "bob"
	e2.Age = 12
	t.Log(e)     // {1 Edge 10}
	t.Log(e1)    // {2 Plan 11}
	t.Log(e2)    // &{2 bob 12}
	t.Log(e.Id)  // 1
	t.Log(e1.Id) // 2
	t.Log(e2.Id) // 3
	t.Logf("e is %T", e)
	t.Logf("e2 is %T", e2)
	t.Logf("e is %T", &e)
}

func TestStructOperations(t *testing.T) {
	e := Employee{1, "Bob", 20}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
}
