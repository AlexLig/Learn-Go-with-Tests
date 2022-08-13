package integers

func SumAllTails(numberCollections ...[]int) []int {
	tailsSum := []int{}
	for _, numberCollection := range numberCollections {
		tails := []int{}

		if len(numberCollection) > 0 {
			tails = numberCollection[1:]
		}

		sum := Sum(tails)
		tailsSum = append(tailsSum, sum)
	}
	return tailsSum
}
