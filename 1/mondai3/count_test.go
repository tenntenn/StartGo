package count

import (
	"testing"
)

func TestCount(t *testing.T) {
	n := 10

	j := 0
	for i := range Count(n) {
		if i != j {
			t.Errorf("i and j must be equal.")
		}
		j++
	}

	if j != n + 1 {
		t.Errorf("j must be %d.", n + 1)
	}
}

func TestStepBy(t *testing.T) {
	step := 3
	n := 200
	j := 0
	for i := range StepBy(Count(n), step) {
		if i != j {
			t.Errorf("i and j must be equal.")
		}
		j += step
	}

	if j != 201 {
		t.Errorf("j must be %d.", 201)
	}
}
