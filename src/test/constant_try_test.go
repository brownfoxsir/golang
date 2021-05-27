package test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)
const (
	Readable = 1 << iota // 0001
	//iota计数器，初始化为0 iota在const关键字出现时将被重置为0(const内部的第一行之前)，const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)
	Writable // 2 0010
	Executable // 4 0100
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}

func TestConstantTry1(t *testing.T) {
	a := 1 // 0001
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
