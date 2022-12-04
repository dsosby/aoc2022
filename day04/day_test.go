package main

import (
	"testing"
)

func TestProblem1(t *testing.T) {
	solution := Problem1()
	if solution != 441 {
		t.Log("Expected 441")
		t.Fail()
	}
}

func TestProblem2(t *testing.T) {
	solution := Problem2()
	if solution != 861 {
		t.Log("Expected 861")
		t.Fail()
	}
}
