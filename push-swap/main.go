package main

import (
	"os"
	"strings"
	"strconv"
	"fmt"
	"pile"
)

func Main() {
	entre := os.Args[1]
	pile_str := strings.Split(string(entre)," ")
	fmt.Println("yes")

	//erreur

	var pile_a []int
	for _, val := range pile_str {
		val_int,er := strconv.Atoi(val)
		if er != nil {
			fmt.Println("Error")
			return
		}
		pile_a = append(pile_a,val_int)
	}

	//var pile_b []int

	if pile_a[0] > pile_a[1] {
		fmt.Println("sa")
		pile_a = pile.Sa(pile_a)
	}
	
}