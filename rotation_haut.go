package pile

func Ra(a []int) []int {
	val := a[0]
	a = a[1:]
	a = append(a, val)
	return a
}

func Rb(b []int) []int {
	val := b[0]
	b = b[1:]
	b = append(b, val)
	return b
}

func Rr(a []int,b []int) ([]int,[]int) {
	a = Ra(a)
	b = Rb(b)
	return a,b
}