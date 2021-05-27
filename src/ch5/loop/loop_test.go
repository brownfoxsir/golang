package loop

import "testing"

func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}
}

func TestInfiniteLoop(t *testing.T){
	n := 0
	for {
		t.Log(n)
		n++
	}
}

func TestMissingNumber(t *testing.T) {
	nums := [...]int{0,1,3,4}
	for key, value:= range nums {
		if key ^ value !=0 {
			t.Log(key)
		}
	}
}

func missingNumber(nums []int) int {
	for key, value:= range nums {
		if key ^ value !=0 {
			return key
		}
	}
	return 0
}