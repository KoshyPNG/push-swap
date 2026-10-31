package main

import (
	"os"
	"strings"
	"strconv"
	"fmt"
	"pile"
)

func main() {
	entre := os.Args[1]
	pile_str := strings.Split(string(entre)," ")

	var pile_a []int
	for _, val := range pile_str {
		val_int,er := strconv.Atoi(val)
		if er != nil {
			fmt.Println("Error")
			return
		}
		pile_a = append(pile_a,val_int)
	}

	if pile.Dupli(pile_a) {
		fmt.Println("Error doublon")
		return
	}
	
	var pile_b []int

	if len(pile_a) < 2 {
		return
	}else if len(pile_a) == 2 {

		if pile_a[0] > pile_a[1] {
			fmt.Println("sa")
			pile_a = pile.Sa(pile_a)
		}
		return

	} else if len(pile_a) == 3 {
		pile_a = pile.Algo_trois(pile_a)
		return

	} else if len(pile_a) < 9 {
		pile_a = pile.Algo_cinq(pile_a,pile_b)

	} else if len(pile_a) <= 1023 {
		pile_a = pile.Algo_binaire(pile_a,pile_b)
	} else {
		fmt.Println("pile_a trop grande !!!!!")
	}
}