package main

import (
	"fmt"
	"test-with-go/math"
)

func main() {
	sum := math.Sum([]int{10, -2, 3})
	if sum != 3 {
		msg := fmt.Sprintf("FAIL : Wanted 11 but returned %d", sum)
		panic(msg)
	}
	add := math.Add(5, 10)
	if add != 15 {
		msg := fmt.Sprintf("FAIL : Wanted 15 but returned %d", add)
		panic(msg)
	}
	fmt.Println("PASS")
}
