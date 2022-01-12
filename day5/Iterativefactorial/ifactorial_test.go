package Iterativefactorial

import "testing"

func TestIfactorial(t *testing.T) {
	//table driven tests
	//slice literals

	s := []struct {
		input  int
		expect int
	}{
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
	}

	for _, v := range s {
		output := Ifactorial(v.input)
		expected := v.expect
		if output != expected {

			t.Errorf("expected is %v and output got is %v", expected, output)

		}
	}

}
