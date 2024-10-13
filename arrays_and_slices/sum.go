package main

func Sum(numbers []int) (result int) {
	for _, v := range numbers {
		result += v
	}
	return result
}

func SumAllTails(numberSlices ...[]int) (result []int) {

	for _, v := range numberSlices {
		// check len of slice
		if len(v) == 0 {
			result = append(result, 0)
		} else {
			// strip the head of slice
			result = append(result, Sum(v[1:]))
		}
	}
	return
}
