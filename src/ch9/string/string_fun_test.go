package string

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFun(t *testing.T) {
	s := "a,b,c,d"
	parts := strings.Split(s, ",")
	t.Log(parts)
	for i, v := range parts {
		t.Log(i, v)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestConvert(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str" + s)
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}

}
