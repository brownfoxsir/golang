package _map

import (
	"testing"
)

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])                 // 4
	t.Logf("len m1=%d", len(m1)) // len m1=3
	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d", len(m2)) // len m2=1
	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3)) // len m3=0
}

func TestAccessNotExistKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1]) // 0
	m1[2] = 1
	t.Log(m1[2])
	if v, ok := m1[3]; ok {
		t.Logf("Key 3's value is %d", v)
	} else {
		t.Log(v, ok)                    // 0 false
		t.Logf("Key 3 is not existing") // Key 3 is not existing
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for key, value := range m1 {
		t.Log(key, value)
	}
}
