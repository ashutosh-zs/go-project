package leapyear

import "testing"

func TestLeapYear(t *testing.T) {
	//slice literals

	s := []struct {
		inp    int
		expect string
	}{
		{2001, "Leap year"},
		{1900, "Not a leap year"},
		{2004, "Leap year"},
		{2005, "Not a leap year"},
	}

	for _, v := range s {
		output := checkLeap(v.inp)
		expected := v.expect
		if output != expected {

			t.Errorf("expected is %v and output got is %v", expected, output)

		}
	}

}
