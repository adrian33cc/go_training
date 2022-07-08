package expenses

func Average(expensesList ...float64) float64 {

	return SumTotal(expensesList...) / float64(len(expensesList))

}

func SumTotal(expensesList ...float64) float64 {
	var sum float64

	for _, expense := range expensesList {
		sum += expense
	}

	return sum
}

func Max(expensesList ...float64) float64 {
	var max float64

	if len(expensesList) == 0 {
		return 0
	}

	for _, item := range expensesList {
		if item > max {
			max = item
		}
	}
	return max
}

func Min(expensesList ...float64) float64 {
	min := expensesList[0]

	if len(expensesList) == 0 {
		return 0
	}

	for _, item := range expensesList {
		if item < min {
			min = item
		}
	}
	return min
}
