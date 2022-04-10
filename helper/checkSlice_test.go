package helper

import "testing"

func TestCheckSlice(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	if CheckSlice(arr, 10) {
		t.Error("Expected false, got true")
	}
	if !CheckSlice(arr, 1) {
		t.Error("Expected true, got false")
	}
}
