package problem4

import "testing"

func TestFindInclusiveRanges(t *testing.T) {
	if val := FindInclusiveRanges("test.txt"); val != 2 {
		t.Fatalf("Expected 2 got %d", val)
	}
}

func TestFindInclusiveRangesProblem(t *testing.T) {
	if val := FindInclusiveRanges("input.txt"); val != 515 {
		t.Fatalf("Expected 515 got %d", val)
	}
}

func TestFindUnionRanges(t *testing.T) {
	if val := FindUnionRanges("test.txt"); val != 4 {
		t.Fatalf("Expected 4 got %d", val)
	}
}

func TestFindUnionRangesProblem(t *testing.T) {
	if val := FindUnionRanges("input.txt"); val != 883 {
		t.Fatalf("Expected 883 got %d", val)
	}
}
