package arrayslice

func Sum(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum = sum + num
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	sum := []int{}
	for _, nums := range numbersToSum {
		sum = append(sum, Sum(nums))
	}
	return sum
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sum []int
	for _, nums := range numbersToSum {
		if len(nums) == 0 {
			sum = append(sum, 0)
		} else {
			tail := nums[1:]
			sum = append(sum, Sum(tail))
		}
	}

	return sum
}
