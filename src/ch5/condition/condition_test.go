package condition

import (
	"fmt"
	"runtime"
	"testing"
)

func TestIfMultiSec(t *testing.T) {
	if a := 1 == 1; a {
		t.Log(a)
	}

	//if value, err := someFunc(); err == nil {
	//	t.Log(value)
	//}
}

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even", i)
		case 1, 3:
			t.Log("Odd", i)
		default:
			t.Log("It is not 0-3", i)

		}
	}
}

func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even", i)
		case i%2 == 1:
			t.Log("Odd", i)
		default:
			t.Log("Unknown", i)

		}
	}
}

func TestSwitchCase(t *testing.T) {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
		//break
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}
