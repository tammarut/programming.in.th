package main

import "testing"

func TestGetGradeWhenPassSum80ShouldReturnA(t *testing.T) {
	result := getGrade(80)
	if result != "A" {
		t.Error("Expected: A", "but Got->", result)
	}
}
func TestGetGradeWhenPassSum75ShouldReturnBplus(t *testing.T) {
	result := getGrade(75)
	if result != "B+" {
		t.Error("Expected: B+", "but Got->", result)
	}
}
func TestGetGradeWhenPassSum70ShouldReturnB(t *testing.T) {
	result := getGrade(70)
	if result != "B" {
		t.Error("Expected: B", "but Got->", result)
	}
}
func TestGetGradeWhenPassSum65ShouldReturnCplus(t *testing.T) {
	result := getGrade(65)
	if result != "C+" {
		t.Error("Expected: C+", "but Got->", result)
	}
}

func TestGetGradeWhenPassSum60ShouldReturnC(t *testing.T) {
	result := getGrade(60)
	if result != "C" {
		t.Error("Expected: C", "but Got->", result)
	}
}

func TestGetGradeWhenPassSum55ShouldReturnDplus(t *testing.T) {
	result := getGrade(55)
	if result != "D+" {
		t.Error("Expected: D+", "but Got->", result)
	}
}

func TestGetGradeWhenPassSum50ShouldReturnD(t *testing.T) {
	result := getGrade(50)
	if result != "D" {
		t.Error("Expected: D", "but Got->", result)
	}
}

func TestGetGradeWhenPassSum49ShouldReturnF(t *testing.T) {
	result := getGrade(49)
	if result != "F" {
		t.Error("Expected: D", "but Got->", result)
	}
}

func TestGetGrade_WhenPassTableScore_ShouldReturnGrade(t *testing.T) {
	var mocks = []struct {
		score int
		grade string
	}{
		{80, "A"},
		{75, "B+"},
		{70, "B"},
		{65, "C+"},
		{60, "C"},
		{55, "D+"},
		{54, "D"},
		{49, "F"},
	}

	for _, v := range mocks {
		result := getGrade(v.score)
		if result != v.grade {
			t.Errorf("Expected: %s but Got: %s", v.grade, result)
		}
	}
}
