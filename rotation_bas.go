package pile

import "fmt"

func Rra(a []int) []int {
	if !vrai {
		fmt.Println("rra")
	}
	val := a[len(a)-1]
	var new []int
	new = append(new, val)
	for _, v := range a[:len(a)-1] {
		new = append(new, v)
	}
	return new
}

func Rrb(b []int) []int {
	if !vrai {
		fmt.Println("rrb")
	}
	val := b[len(b)-1]
	var new []int
	new = append(new, val)
	for _, v := range b[:len(b)-1] {
		new = append(new, v)
	}
	return new
}

func Rrr(a []int, b []int) ([]int, []int) {
	vrai = true
	fmt.Println("rrr")
	a = Rra(a)
	b = Rrb(b)
	vrai = false
	return a, b
}
