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

func TestSumAll(t *testing.T) {
	foo := []int{1, 2, 3}
	bar := []int{5, 6}

	got := SumAll(foo, bar)
	want := []int{6, 11}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d", got, want)
	}
}

func ExampleSumAll() {
	foo := []int{1, 2}
	bar := []int{5, 6}
	fmt.Println(SumAll(foo, bar))

	// Output: [3 11]
}
