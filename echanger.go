package pile

func Sa(a []int) []int {
	a[0],a[1] = a[1],a[0]
	return a
}

func Sb(b []int) []int {
	b[0],b[1] = b[1],b[0]
	return b
}

func Ss(a []int,b []int) ([]int,[]int) {
	a  = Sa(a)
	b = Sb(b)
	return a,b
}