package bench

import "sync"

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

// FibMemo - not threadsafe
func FibMemo(n int) int {
	if len(memo) <= n {
		FibMemo(n - 1)                                 // make sure the memo is filled in up to n-1
		memo = append(memo, FibMemo(n-1)+FibMemo(n-2)) // append the nth value
	}
	return memo[n]
}

var (
	mutex   sync.RWMutex
	memoMap = map[int]int{
		0: 0,
		1: 1,
	}
)

// FibMemoThreadSafe
func FibMemoThreadSafe(n int) int {

	// lock for reading
	mutex.RLock()

	// if value exist in the map
	if v, ok := memoMap[n]; ok {
		defer mutex.RUnlock()
		return v
	}

	// unlock for reading
	mutex.RUnlock()

	n2 := FibMemoThreadSafe(n - 2)
	n1 := FibMemoThreadSafe(n - 1)

	// lock for writing
	mutex.Lock()

	memoMap[n] = n1 + n2

	// unlock for writing
	mutex.Unlock()

	return n1 + n2
}
