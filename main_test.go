package main

import (
	"reflect"
	"testing"
)

func TestIsEven(t *testing.T) {
	assertOdd := isEven(9)
	assertEven := isEven(10)

	if assertOdd != false {
		t.Errorf("Expected false, but received %v", assertOdd)
	}

	if assertEven != true {
		t.Errorf("Expected true, but received %v", assertEven)
	}
}

func TestInitNumSlice(t *testing.T) {
	numSlice0 := initNumSlice(0)
	numSlice10 := initNumSlice(10)
	assert0 := []int{}
	assert10 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	if len(numSlice10) != len(assert10) {
		t.Errorf("Expected %v, but received %v", assert10, numSlice10)
	}
	
	if reflect.DeepEqual(numSlice10, assert10) != true {
		t.Errorf("Expected %v, but received %v", assert10, numSlice10)
	}

	if len(numSlice0) != len(assert0) {
		t.Errorf("Expected %v, but received %v", assert0, numSlice0)
	}
}