package fibo

import (
	"testing"
)

func TestFibo(t *testing.T) {
	expect := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987}
	fibo := Fibo()
	i := 0
	for v := range fibo{
		if i >= len(expect) {
			break
		}
		if v != expect[i] {
			t.Errorf("v and expect[i] must be equal.")
		}

		i++
	}
}
