package main

import (
	"testing"
)

func Test_binary_search(t *testing.T) {
	array := []int64{1, 3, 4, 5}
	to_search1 := 3
	if !binary_search(array, int64(to_search1)) {
		t.Errorf("expected true, got false")
	}

	to_search2 := 6
	if binary_search(array, int64(to_search2)) {
		t.Errorf("expected false, go true")
	}
}
