package pile

func Trier(a []int) bool {
	if len(a) < 2 {
		return true
	}
	for i := 1; i < len(a); i++ {
		if a[i-1] > a[i] {
			return false
		}
	}
	return true
}