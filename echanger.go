package pile

import "fmt"

var vrai bool = false

func Sa(a []int) []int {
	if !vrai {
		fmt.Println("sa")
	}
	
	a[0],a[1] = a[1],a[0]
	return a
}

func Sb(b []int) []int {
	if !vrai {
		fmt.Println("sb")
	}
	
	b[0],b[1] = b[1],b[0]
	return b
}

func Ss(a []int,b []int) ([]int,[]int) {
	vrai = true
	fmt.Println("ss")
	a  = Sa(a)
	b = Sb(b)
	vrai = false
	return a,b
}