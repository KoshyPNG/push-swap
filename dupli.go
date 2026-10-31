package pile

func Dupli(a []int) bool {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] == a[j] {
				return true
			}
		}
	}
	return false
}