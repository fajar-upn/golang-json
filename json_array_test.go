package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

/**
 * JSON array
 * other the data type in the form of object, commonly in the form of JSON
 * we sometimes use data type n Array
 *
 * Array in JSON likes with Array in javascrpt, that's array can fullfill
 * primitive data type, or complex data type (Object or Arary)
 *
 * in Go-lang data type represnt in slice
 *
 * Convertion from JSOn to JSON do automatically with package json
 * use slice data type
 */

type Address struct {
	Street  string
	Country string
}

type Customer1 struct {
	FirstName, LastName string
	Age                 int
	Married             bool
	Hobbies             []string
	Adresses            []Address
}

func TestArrayEncodeJSON(t *testing.T) {
	customer := Customer1{
		FirstName: "Fajar",
		LastName:  "Sidiq",
		Hobbies:   []string{"Badminton", "Playing Games"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestArrayDecodeJSON(t *testing.T) {
	jsonString := `{"FirstName": "Fajar", "LastName": "sidiq", "Age": 23, "Married":false, "Hobbies":["Badminton", "Gaming"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
}

func TestJSONArrayComplexEncode(t *testing.T) {
	customer := Customer1{
		FirstName: "Fajar",
		Age:       24,
		Hobbies:   []string{"Badminton", "Gaming"},
		Adresses: []Address{
			{
				Street:  "Jl. Solo",
				Country: "Indonesia",
			},
			{
				Street:  "Jl. Madiun",
				Country: "Indonesia",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Fajar","LastName":"","Age":24,"Married":false,"Hobbies":["Badminton","Gaming"],"Adresses":[{"Street":"Jl. Solo","Country":"Indonesia"},{"Street":"Jl. Madiun","Country":"Indonesia"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer1{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
}
