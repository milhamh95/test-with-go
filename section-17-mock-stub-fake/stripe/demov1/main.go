package demo

import (
	"encoding/json"
	"fmt"
	"os"
	envjson "test-with-go/envjson"
	stripe "test-with-go/section-17-mock-stub-fake/stripe/v1"
)

func main() {

	err := envjson.Setup()
	if err != nil {
		panic(err)
	}
	c := stripe.Client{
		Key: os.Getenv("KEY"),
	}
	charge, err := c.Charge(2000, "tok_mastercard", "Charge for demo purposes")
	if err != nil {
		panic(err)
	}

	fmt.Println(charge)

	jsonBytes, err := json.Marshal(charge)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonBytes))

}
