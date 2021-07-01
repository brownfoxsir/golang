package _interface

import (
	"fmt"
	"testing"
	"time"
)

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}

type IntConv func(oop int) int

func timeSpent(inner IntConv) IntConv{
	return func(n int) int {
		start := time.Now()
		ret := inner(n)

		fmt.Println("time spent: ", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(ops int) int {
	time.Sleep(time.Second * 1)
	fmt.Println("hello slow fun")
	return ops
}

func TestFun(t *testing.T) {
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))

}