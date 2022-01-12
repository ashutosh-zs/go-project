package PascalTriangle

func PascalTriangle(rows int) [][]int {

	res := make([][]int, rows)

	for i := 0; i < rows; i++ {
		res[i] = make([]int, i+1)
		res[i][0], res[i][i] = 1, 1

		for j := 0; j < i; j++ {
			res[i][j] = res[i-1][j] + res[i-1][j-1]
		}

	}

	return res

}
