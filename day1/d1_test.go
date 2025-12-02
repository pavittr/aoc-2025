package day1

import (
	"fmt"
	"os"
	"testing"
)

func TestSolve(t *testing.T) {
	puzzle := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
	result := Solve(puzzle)
	expected := 3

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestFile(t *testing.T) {
	// Read the test case from testcases/day1.test
	// and the answer from day1.answer
	testCase, err := os.ReadFile("../testcases/day1.test")
	if err != nil {
		t.Fatalf("Failed to read test case: %v", err)
	}

	answerData, err := os.ReadFile("../testcases/day1.answer")
	if err != nil {
		t.Fatalf("Failed to read answer file: %v", err)
	}

	var expected int
	_, err = fmt.Sscanf(string(answerData), "%d", &expected)
	if err != nil {
		t.Fatalf("Failed to parse expected answer: %v", err)
	}

	result := Solve(string(testCase))

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}
