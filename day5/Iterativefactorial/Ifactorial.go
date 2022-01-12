package Iterativefactorial

func Ifactorial(n int) int {
	fact := 1
	for i := 1; i <= n; i++ {
		fact = fact * i
	}
	return fact
}
