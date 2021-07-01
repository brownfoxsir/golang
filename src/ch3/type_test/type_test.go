package type_test

import (
	"math"
	"testing"
)

type myInt int64

func TestImplicit(t *testing.T) {
	var a = 1
	var b int64
	b = int64(a)

	var c myInt
	c = myInt(b)
	var d = math.MaxInt8
	t.Log(a, b, c, d)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	//aPtr += 1  invalid operation: aPtr += 1 (mismatched types *int and int
	b := *aPtr + 1            //取指针所指向地址存储的值
	t.Log(a, aPtr, b)         // 1 0xc0000a6198 2
	t.Logf("%T, %T", a, aPtr) // int, *int
}

func TestString(t *testing.T) {
	var s string
	t.Log("==" + s + "==") //====
	t.Log(len(s))
	if s == "" {
		t.Log("空")
	}
}
