package day1

import (
	"reflect"
	"testing"
)

func TestMatch(t *testing.T) {
	tests := []struct {
		searchString string
		findString   string
		indices      []int
	}{
		{"onetwothree", "one", []int{0}},
		{"onetwothree", "two", []int{3}},
		{"onetwothree", "three", []int{6}},
		{"oneightasd", "eight", []int{2}},
		{"asdlkje", "one", []int{}},
		{"", "", []int{}},
		{"onetwoone", "one", []int{0, 6}},
	}

	for i, tt := range tests {
		foundIndices := FindStringIndex(tt.searchString, tt.findString)
		if !reflect.DeepEqual(foundIndices, tt.indices) {
			t.Errorf("test[%d] - findStringIndex did not return correct index. got=%v, expected=%v", i, foundIndices, tt.indices)
		}
	}

}
