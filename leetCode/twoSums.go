package leetcode

func twoSum(numbers []int, target int) []int {
	complementsMap := make(map[int]int)

	for index, number := range numbers {
		differenceFromTarget := target - number
		complementIndex, ok := complementsMap[number]
		if ok {
			return []int{complementIndex, index}

		}
		complementsMap[differenceFromTarget] = index

	}

	return []int{0, 0}
}
