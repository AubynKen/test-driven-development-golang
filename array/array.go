package array

func Sum(numbers []int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}

func SumAll(numberSlices ...[]int) []int {
	var res []int
	for _, numbers := range numberSlices {
		res = append(res, Sum(numbers))
	}
	return res
}

func SumAllTails(numberSlices ...[]int) []int {
	var res []int
	for _, numbers := range numberSlices {
		if len(numbers) == 0 {
			res = append(res, 0)
			continue
		}
		res = append(res, Sum(numbers[1:]))
	}
	return res
}
