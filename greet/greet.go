package greet

import (
	"fmt"
)

func Hello(name string) (string, error) {
	return fmt.Sprintf("Hello, %s", name), nil
}

func Page(checkIns map[string]bool) {
	for name, checkIn := range checkIns {
		if !checkIn {
			fmt.Printf("Paging %s: please see the front desk to check in", name)
		}
	}
}
