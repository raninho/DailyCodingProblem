package problem1

import "testing"

func equal(a, b []int) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestProblem1Scene1(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	output := []int{1, 5, 2, 4, 3}

	result := solution(input)

	if !equal(output, result) {
		t.Fatal()
	}
}

func TestProblem1Scene2(t *testing.T) {
	input := []int{1, 2, 3, 4}
	output := []int{1, 4, 2, 3}

	result := solution(input)

	if !equal(output, result) {
		t.Fatal()
	}
}

func TestProblem1Scene3(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7}
	output := []int{1, 7, 2, 6, 3, 5, 4}

	result := solution(input)

	if !equal(output, result) {
		t.Fatal()
	}
}
