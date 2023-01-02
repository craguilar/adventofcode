package problem3

import "testing"

func TestFindReorganizationCorrectness(t *testing.T) {
	if val := FindReorganization("test.txt"); val != 157 {
		t.Errorf("Expected 157 got %d", val)
	}
}

func TestProblem(t *testing.T) {
	if val := FindReorganization("input.txt"); val != 8053 {
		t.Errorf("Expected 157 got %d", val)
	}
}

func TestFindBadgeCorrectness(t *testing.T) {
	if val := FindBadge("test.txt"); val != 70 {
		t.Errorf("Expected 70 got %d", val)
	}
}

func TestFindBadgeProblem(t *testing.T) {
	if val := FindBadge("input.txt"); val != 2425 {
		t.Errorf("Expected 2425 got %d", val)
	}
}
