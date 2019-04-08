package main

import "testing"

func TestGetGradeWhenPassSumShouldReturnString(t *testing.T) {
	result := getGrade(80)
	if result != "A" {
		t.Error("Expected: A", "but Got->", result)
	}
}
