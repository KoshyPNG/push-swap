package pile

import "fmt"

func Ra(a []int) []int {
	if !vrai {
		fmt.Println("ra")
	}
	val := a[0]
	a = a[1:]
	a = append(a, val)
	return a
}

func Rb(b []int) []int {
	if !vrai {
		fmt.Println("rb")
	}
	val := b[0]
	b = b[1:]
	b = append(b, val)
	return b
}

func Rr(a []int, b []int) ([]int, []int) {
	vrai = true
	fmt.Println("rr")
	a = Ra(a)
	b = Rb(b)
	vrai = false
	return a, b
}
