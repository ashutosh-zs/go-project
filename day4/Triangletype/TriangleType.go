package Triangletype

func checkTypeTraingle(x, y, z int) string {
	if x < 0 || y < 0 || z < 0 || x+y < z || x+z < y || y+z < x {
		return "Not A traingle"
	} else if x == y && y == z {
		return "Equilatral triangle"
	} else if x != y && y != z {
		return "Scalene Triangle"
	}

	return "Isosceles Traingle"
	// if x+y ==z || x+z == y || y+z == x
	//return "Degenerate Tr

}
