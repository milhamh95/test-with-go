package math

func Sum(numbers []int) int {
	sum := 0
	for n := range numbers {
		sum += n
	}
	return sum
}

func Add(a, b int) int {
	return a + b
}
