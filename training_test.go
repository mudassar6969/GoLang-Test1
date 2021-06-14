package main

import (
	"fmt"
	"testing"
)

func TestFindWordsCountBasic(t *testing.T) {
	count := findWordsCount("saveChangesInTheEditor")
	if count != 5 {
		t.Error("count should be 5 instead of", count)
	}
}

type WordsTest struct {
	Word  string
	Count int
}

func TestFindWordsCountTableDriven(t *testing.T) {
	tests := []WordsTest{{"testDrivenDevelopment", 3}, {"testAgain", 2}, {"thisIsSimpleAgain", 4}}

	for _, item := range tests {

		testName := fmt.Sprintf("%s, %d", item.Word, item.Count)
		t.Run(testName, func(t *testing.T) {
			count := findWordsCount(item.Word)
			if item.Count != count {
				t.Error("Count is not correct", count)
			}
		})
	}
}
