package _func

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValue() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(oop int) int) func(oop int) int {
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
	a, _ := returnMultiValue()
	t.Log(a)
	//t.Log(_) 会报错，不能这么写，'_'表示忽略这个返回的变量
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))

}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 4, 5))
}

func Clear() {
	fmt.Println("clear resources")
}

func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("start")
	panic("error")
	//fmt.Println("不会执行这条语句")
}
