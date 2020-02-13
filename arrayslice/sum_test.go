package arrayslice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func ExampleSum() {
	foo := []int{1, 2}
	fmt.Println(Sum(foo))
	// Output: 3
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("make the sum of some slices", func(t *testing.T) {
		foo := []int{1, 2, 3}
		bar := []int{5, 6}

		got := SumAllTails(foo, bar)
		want := []int{5, 6}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		foo := []int{}
		bar := []int{2, 1, 3}

		got := SumAllTails(foo, bar)
		want := []int{0, 4}
		checkSums(t, got, want)
	})
}

func ExampleSumAllTails() {
	foo := []int{1, 2}
	bar := []int{5, 6}
	fmt.Println(SumAllTails(foo, bar))

	// Output: [2 6]
}
