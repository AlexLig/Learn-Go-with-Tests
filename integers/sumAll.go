package integers

func SumAll(numbersToSum ...[]int) []int {
	result := []int{}
	for _, numberCollection := range numbersToSum {
		sum := Sum(numberCollection)
		result = append(result, sum)
	}
	return result
}
