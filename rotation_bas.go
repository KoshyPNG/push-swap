package pile

func Rra(a []int) []int {
	val := a[len(a)]
	var new []int
	new = append(new, val) 
	for _, v := range a[:len(a)] {
		new = append(new, v) 
	}
	return new
}

func Rrb(b []int) []int {
	val := b[len(b)]
	var new []int
	new = append(new, val) 
	for _, v := range b[:len(b)] {
		new = append(new, v) 
	}
	return new
}

func Rrr(a []int,b []int) ([]int,[]int) {
	a = Rra(a)
	b = Rrb(b)
	return a,b
}