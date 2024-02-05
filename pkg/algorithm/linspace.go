package algorithm

// create a list of n evenly-spaced integers, starting at min and ending with max
// if max > min, list will be ascending, else it will be descending
func Linspace(min int, max int, n uint32) []int {
	numbers := make([]int, n)
	for i := 0; i < int(n); i++ {
		numbers[i] = min + (i * (max - min) / (int(n) - 1))
	}
	return numbers
}
