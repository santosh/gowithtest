package arrayslice

// Sum takes a slice and returns its sum
func Sum(arr []int) (result int) {
	for _, num := range arr {
		result += num
	}
	return
}

// SumAll takes multiple slice and returns a single slice.
// Each index representing sum of each slice passed.
func SumAll(numToSum ...[]int) (sums []int) {
	for _, numbers := range numToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}
