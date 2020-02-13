package arrayslice

// Sum takes a slice and returns its sum
func Sum(arr []int) (result int) {
	for _, num := range arr {
		result += num
	}
	return
}
