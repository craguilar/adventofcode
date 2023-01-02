package problem5

import "testing"

func TestCraneSimulationCorrecntess(t *testing.T) {
	if value := CraneSimulation("test.txt", false); value != "CMZ" {
		t.Fatalf("Expected CMZ got %s", value)
	}
}

func TestCraneSimulationProblem(t *testing.T) {
	if value := CraneSimulation("input.txt", false); value != "WCZTHTMPS" {
		t.Fatalf("Expected WCZTHTMPS got %s", value)
	}
}

func TestCraneSimulationMultiPickCorrecntess(t *testing.T) {
	if value := CraneSimulation("test.txt", true); value != "MCD" {
		t.Fatalf("Expected MCD got %s", value)
	}
}

func TestCraneSimulationMultiPickProblem(t *testing.T) {
	if value := CraneSimulation("input.txt", true); value != "BLSGJSDTS" {
		t.Fatalf("Expected BLSGJSDTS got %s", value)
	}
}
