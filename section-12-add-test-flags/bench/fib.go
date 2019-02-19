package bench

// FibRecursive - fibonacci in recursive
func FibRecursive(n int) int {
	if n < 2 {
		return n
	}
	return FibRecursive(n-1) + FibRecursive(n-2)
}

// FibIterative - fibonacci in interative
func FibIterative(n int) int {
	if n < 2 {
		return n
	}
	a, b := 0, 1

	for i := 1; i < n; i++ {
		a, b = b, a+b
	}

	return b

}

var memo = []int{0, 1}

func FibMemo(n int) int {
	if len(memo) <= n {
		FibMemo(n - 1)                                 // make sure the memo is filled in up to n-1
		memo = append(memo, FibMemo(n-1)+FibMemo(n-2)) // append the nth value
	}
	return memo[n]
}
