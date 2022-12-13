package problem2

import "testing"

func TestCorrectness(t *testing.T) {
	if val := Strategy("test.txt"); val != 15 {
		t.Errorf("Expected 15 got %d", val)
	}
}

func TestProblem(t *testing.T) {
	val := Strategy("input.txt")
	t.Logf("Got %d", val)
}

func TestProblemOnPlayCalculation(t *testing.T) {
	val := CalculateStrategy("input.txt")
	t.Logf("Got %d", val)
}

func TestCorrectnessOnPlayCalculation(t *testing.T) {
	if val := CalculateStrategy("test.txt"); val != 12 {
		t.Errorf("Expected 12 got %d", val)
	}
}
