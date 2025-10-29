package pacote

func Sum(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

func Somar(nums ...int) int {
	var out int
	for _, n := range nums {
		out += n
	}
	return out
}
