package biz

func Sum(array []int) int {
	sum := 0
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	return sum
}
