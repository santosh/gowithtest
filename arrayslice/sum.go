package arrayslice

// Sum takes an array of integers and returns its sum
func Sum(arr [5]int) (result int) {
	for i := 0; i < len(arr); i++ {
		result += arr[i]
	}
	return
}
