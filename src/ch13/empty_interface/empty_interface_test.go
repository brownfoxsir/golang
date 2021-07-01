package empty_interface

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}){
	switch v:=p.(type) {
	case int:
		fmt.Println("int",v)
	case string:
		fmt.Println("str",v)
	default:
		fmt.Println("unknown type")

	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
}