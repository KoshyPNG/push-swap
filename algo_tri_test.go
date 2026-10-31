package pile

import (
	"reflect"
	"sort"
	"testing"
)

func assertSortedValues(t *testing.T, got, input []int) {
	t.Helper()

	want := append([]int(nil), input...)
	sort.Ints(want)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("tri obtenu %v, attendu %v", got, want)
	}
}

func TestAlgoTrois(t *testing.T) {
	tests := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}

	for _, input := range tests {
		t.Run("", func(t *testing.T) {
			got := Algo_trois(append([]int(nil), input...))
			assertSortedValues(t, got, input)
		})
	}
}

func TestAlgoCinq(t *testing.T) {
	input := []int{8, 3, 6, 1, 5, 2, 7, 4}
	got := Algo_cinq(append([]int(nil), input...), nil)
	assertSortedValues(t, got, input)
}

func TestAlgoBinaire(t *testing.T) {
	input := make([]int, 100)
	for i := range input {
		input[i] = len(input) - i - 1
	}

	got := Algo_binaire(append([]int(nil), input...), nil)
	assertSortedValues(t, got, input)
}

func TestTrier(t *testing.T) {
	tests := []struct {
		name string
		list []int
		want bool
	}{
		{name: "empty", list: nil, want: true},
		{name: "single value", list: []int{4}, want: true},
		{name: "sorted", list: []int{1, 2, 3}, want: true},
		{name: "unsorted", list: []int{1, 3, 2}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Trier(test.list); got != test.want {
				t.Fatalf("Trier(%v) = %t, attendu %t", test.list, got, test.want)
			}
		})
	}
}
