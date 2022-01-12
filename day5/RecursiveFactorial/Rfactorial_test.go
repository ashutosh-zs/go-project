package RecursiveFactorial

import "testing"

func TestRecursiveFactorial(t *testing.T) {
	s := []struct {
		input  int
		expect int
	}{
		{1, 1},
		{2, 2},
		{3, 6},
	}

	for _, v := range s {
		output := RecursiveFactorial(v.input)
		expected := v.expect

		if output != expected {
			t.Errorf("Expected is %v Output is %v", expected, output)
		}
	}
}
