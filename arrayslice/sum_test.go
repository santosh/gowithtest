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
	foo := []int{1, 2, 3}
	bar := []int{5, 6}

	got := SumAllTails(foo, bar)
	want := []int{5, 6}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d", got, want)
	}
}

func ExampleSumAllTails() {
	foo := []int{1, 2}
	bar := []int{5, 6}
	fmt.Println(SumAllTails(foo, bar))

	// Output: [2 6]
}
