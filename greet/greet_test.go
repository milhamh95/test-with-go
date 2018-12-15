package greet

import "fmt"

// func ExampleHello() {
// 	greeting, err := Hello("Jon")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(greeting)

// 	// Output:
// 	// Hello, Jon
// }

// func ExampleDemo() {
// 	greeting, err := Hello("Jon")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(greeting)

// 	// Output:
// 	// Hello, Jon
// }

// func ExampleDemo_Hello() {
// 	greeting, err := Hello("Jon")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(greeting)

// 	// Output:
// 	// Hello, Jon
// }

func ExampleHello_spanish() {
	greeting, err := Hello("Jon")
	if err != nil {
		panic(err)
	}
	fmt.Println(greeting)

	// Output:
	// Hello, Jon
}
