package client

import "testing"
import "golang/src/ch15/series"

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacciSeries(5))
	t.Log(series.Square(5))
}
