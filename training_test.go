package main

import "testing"

func TestFindWordsCount(t *testing.T) {
	count := findWordsCount("saveChangesInTheEditor")
	if count != 5 {
		t.Error("count should be 5 instead of", count)
	}
}
