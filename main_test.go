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
	numSlice5 := initNumSlice(5)
	assert0 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert5 := []int{1, 2, 3, 4, 5}

	if len(numSlice5) != len(assert5) {
		t.Errorf("Expected %v, but received %v", assert5, numSlice5)
	}
	
	if reflect.DeepEqual(numSlice5, assert5) != true {
		t.Errorf("Expected %v, but received %v", assert5, numSlice5)
	}

	if len(numSlice0) != len(assert0) {
		t.Errorf("Expected %v, but received %v", assert0, numSlice0)
	}
}