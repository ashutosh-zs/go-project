package PascalTriangle

import (
	"reflect"
	"testing"
)

func TestPascalTriangle(t *testing.T) {
	op := PascalTriangle(3)
	expected := [][]int{
		{1},
		{1, 1},
		{1, 2, 1},
	}
	if reflect.DeepEqual(op, expected) {
		t.Errorf("expected is %v and output got is %v", expected, op)
	}
}
