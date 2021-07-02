package panic_recover

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	defer func() {
		fmt.Println("finally")
	}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
	}()
	defer func() {
		fmt.Println("hello")
	}()
	fmt.Println("Start")
	panic(errors.New("something wrong")) //panic()传递一个空接口，通常会传递一个错误

	//os.Exit(0)
	//fmt.Println("End")
}
