package Triangletype

import "testing"

func TestTriangleType(t *testing.T) {

	s := []struct {
		x, y, z int

		typee string
	}{
		{1, 2, 3, "Scalene Triangle"},
		{4, 4, 4, "Equilatral triangle"},
		{1, 1, 5, "Not A traingle"},
		{3, 3, 4, "Isosceles Traingle"},
	}

	for _, v := range s {
		output := checkTypeTraingle(v.x, v.y, v.z)
		expected := v.typee
		if output != expected {

			t.Errorf("expected is %v and output got is %v", expected, output)

		}
	}
}
