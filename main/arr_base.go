package main

/* PositiveSum */
func PositiveSum(numbers []int) int {

	if numbers == nil || len(numbers) < 0 {
		return 0
	}
	var sum = 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > 0 {
			sum = sum + numbers[i]
		}
	}
	return sum
}
