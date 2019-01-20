package main

import (
	"fmt"
	"sync"
	"time"
)

var balance = 100

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go spend(30, &wg)
	go spend(40, &wg)
	wg.Wait()
	fmt.Println(balance)
}

func spend(amount int, wg *sync.WaitGroup) {
	fmt.Println("spend")
	fmt.Println(amount)
	b := balance
	time.Sleep(time.Second)
	b -= amount
	fmt.Println("balance in function")
	fmt.Println(b)
	balance = b
	wg.Done()
}
