package packet

func SomeFunc() string {
	return "some string"
}

func Triangle(a, b, c int) int {
	return a + b + c
}

func SomeBenchFunc(n int) {
	sl := make([]int, 0, n)
	for i := 0; i < n; i++ {
		sl = append(sl, i)
	}
	_ = sl
}
