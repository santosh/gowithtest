package arrayslice

// Sum takes a slice and returns its sum
func Sum(arr []int) (result int) {
	for _, num := range arr {
		result += num
	}
	return
}

// SumAllTails takes multiple slice and returns a single slice.
// Each index representing sum of each slice except the first index (head).
func SumAllTails(numToSum ...[]int) (sums []int) {
	for _, numbers := range numToSum {
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}
	return sums
}
