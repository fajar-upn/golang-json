package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

/**
 * Decode JSON
 * Now we already know how to do encode from data type in Go-lang
 * to JSON
 *
 * for do conversion from JSON to data type in Go-lang(Decode).
 * we can use function json.Unmarshal(byte[], interface{})
 *
 * byte[] is a data JSON, whereas interface{} is place for
 * save result convertion, commonly is data type pointer
 */

type Customer struct {
	FirstName, LastName string
	Age                 int
	Married             bool
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName: "Fajar",
		LastName:  "Sidiq",
		Age:       24,
		Married:   false,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestDecodeJSON(t *testing.T) {
	jsonString := `{"FirstName": "Fajar", "Age": 20, "Married": false}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
}
