package greet

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

// func ExampleHello_spanish() {
// 	greeting, err := Hello("Jon")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(greeting)

// 	// Output:
// 	// Hello, Jon
// }

func ExamplePage() {
	checkIns := map[string]bool{
		"Bob":   true,
		"Alice": false,
		"Eve":   false,
		"Janet": true,
		"Susan": true,
		"John":  false,
	}
	Page(checkIns)

	// Unordered output:
	// Paging Alice; please see the front desk to check in.
	// Paging Eve; please see the front desk to check in.
	// Paging John; please see the front desk to check in.
}
