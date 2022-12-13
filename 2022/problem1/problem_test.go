package problem1

import (
	"testing"
)

func TestProblem1(t *testing.T) {
	if TopN(1) != 66616 {
		t.Fail()
	}
}

func TestProblemN(t *testing.T) {
	if TopN(3) != 199172 {
		t.Fail()
	}
}
