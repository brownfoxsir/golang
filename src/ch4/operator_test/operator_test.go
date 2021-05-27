package operator_test

import (
	"testing"
)

const (
	Executable = 1 << iota
	Writable
	Readable
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 4}
	c := [...]int{1, 3, 4, 5}
	//d := [...]int{1, 2}
	t.Log(a == b, a == c) //true, false
	//t.Log(a==d) // a == d (mismatched types [4]int and [2]int)
}

func TestLogicalOperator(t *testing.T) {
	t.Log(true && false)
	t.Log(true || false)
	t.Log(!true)
}
func TestBitClear(t *testing.T) {
	a := 7            // 0111
	a = a &^ Readable // 去掉可读权限 0111 &^ 0100 = 0011
	t.Log("可读", a&Readable == Readable)
	t.Log("可写", a&Writable == Writable)
	t.Log("可执行", a&Executable == Executable)
}
