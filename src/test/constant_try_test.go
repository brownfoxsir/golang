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
	// 每当某个枚举被重置（即后面使用iota重新赋值时），则需要从第一个枚举数到当前的次序, ReStore的次序为3，因此重新赋值为3
	Writable           // 2 0010
	Executable         // 4 0100
	ReStore     = iota // 3
	NewVariable        // 4
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)                 // 1, 2, 3
	t.Log(Writable, Executable, ReStore, NewVariable) //  2 4 3 4
}

func TestConstantTry1(t *testing.T) {
	a := 1 // 0001
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
