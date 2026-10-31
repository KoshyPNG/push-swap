package main

import (
	"bufio"
	"fmt"
	"os"
	"pile"
	"strconv"
	"strings"
)

func main() {
	entre := os.Args[1]
	pile_str := strings.Split(string(entre), " ")

	var pile_a []int
	for _, val := range pile_str {
		val_int, er := strconv.Atoi(val)
		if er != nil {
			fmt.Println("Error")
			return
		}
		pile_a = append(pile_a, val_int)
	}

	if pile.Dupli(pile_a) {
		fmt.Println("Error doublon")
		return
	}
	
	var pile_b []int

	scanner := bufio.NewScanner(os.Stdin)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error")
		return
	}

	for scanner.Scan() {
		inst := scanner.Text()

		switch inst {
		case "sa":
			pile_a = pile.Sa(pile_a)
		case "sb":
			pile_b = pile.Sb(pile_b)
		case "ss":
			pile_a, pile_b = pile.Ss(pile_a, pile_b)
		case "pa":
			pile_a, pile_b = pile.Pa(pile_a, pile_b)
		case "pb":
			pile_a, pile_b = pile.Pb(pile_a, pile_b)
		case "ra":
			pile_a = pile.Ra(pile_a)
		case "rb":
			pile_b = pile.Rb(pile_b)
		case "rr":
			pile_a, pile_b = pile.Rr(pile_a, pile_b)
		case "rra":
			pile_a = pile.Rra(pile_a)
		case "rrb":
			pile_a = pile.Rrb(pile_b)
		case "rrr":
			pile_a, pile_b = pile.Rrr(pile_a, pile_b)
		}
	}
	
	if pile.Trier(pile_a) {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
