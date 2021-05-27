package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 3, 4, 5}
	t.Log(arr[0], arr[1], arr[2]) // 0 0 0
	t.Log(arr1[1])                // 2
	t.Log(arr2)                   // [1 3 4 5]
}

func TestArrayTravel(t *testing.T) {
	arr := [...]int{1, 2, 3, 4}
	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}
	for idx, e := range arr {
		t.Log(idx, e)
	}
}

func TestArraySection(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}
	arrSec := arr[:3]
	arrSec1 := arr[1:]
	arrSec2 := arr[1:2]
	arrSec3 := arr[1:len(arr)]
	t.Log(arrSec)  // [1 2 3]
	t.Log(arrSec1) // [2 3 4 5]
	t.Log(arrSec2) // [2]
	t.Log(arrSec3) //[2 3 4 5]
}
