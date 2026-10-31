package pile

import "fmt"

func Pa(a []int, b []int) ([]int, []int) {
	fmt.Println("pa")
	val := b[0]
	var new []int
	new = append(new, val)
	for _, i := range a {
		new = append(new, i)
	}
	return new, b[1:]
}

func Pb(a []int, b []int) ([]int, []int) {
	fmt.Println("pb")
	val := a[0]
	var new []int
	new = append(new, val)
	for _, i := range b {
		new = append(new, i)
	}
	return a[1:], new
}
