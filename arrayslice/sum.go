package arrayslice

// Sum takes an array of integers and returns its sum
func Sum(arr [5]int) (result int) {
	for _, num := range arr {
		result += num
	}
	return
}
