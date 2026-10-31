package pile


func Algo_trois(a []int) []int {
	if Trier(a) {
		return a
	}
	if a[0] > a[1] && a[0] > a[2] {
		a = Ra(a)
	} else if a[0] < a[1] && a[1] > a[2] {
		a = Rra(a)
	}
	if a[0] > a[1] {
		a = Sa(a)
	}
	return a
}

func Algo_cinq(a []int, b []int) []int {
	if Trier(a) {
		return a
	}
	taille := len(a)
	if taille == 3 {
		a = Algo_trois(a)
		return a
	}
	for len(a) > 3 {
		pos := PositionDuMin(a)
		if pos < (len(a)/2) {
			for pos != 0 {
				a = Ra(a)
				pos--
			}
		} else {
			for pos != 0{
				a = Rra(a)
				pos++
				if pos == len(a) {
					pos = 0
				}
			}
		}
		a,b = Pb(a,b)
	}
	a = Algo_trois(a)
	for len(a) < taille {
		a,b  = Pa(a,b)
	}
	return a
}

func Algo_binaire(a []int, b []int) []int {
	if Trier(a) {
		return a
	}

	n := len(a)
	maxbits := 7
	if n > 127 {
		maxbits = 8
	} else if n > 255 {
		maxbits = 8
	} else if n > 511 {
		maxbits = 9
	} else if  n < 1023 && n > 511 {
		maxbits = 10
	}
	
	for tour := 0; tour < maxbits; tour++ { 
		for k := 0; k < n; k++ {
			num := Chiffre(a[0],tour) 
			if num == 0 {
				a,b = Pb(a,b)
			} else {
				a = Ra(a)
			}
		}
		for len(b) > 0  {
			a,b = Pa(a,b)
		}
	}
	return a
}

func Chiffre(carte, tour int) int {
	return (carte >> tour) & 1
}

func PositionDuMin(a []int) int {
	pos := 0
	for i := 1; i < len(a); i++ {
		if a[i] < a[pos] {
			pos = i
		}
	}
	return pos
}

