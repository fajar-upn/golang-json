package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

/**
 * JSON Object
 * in previous material we do encode data like string, number
 * boolean, and another primitive data
 *
 * although we can do that using encode, because appropriate
 * with interface{} contract, but that is not appropriate with
 * JSON contract
 *
 * If we following json.org contract, the shape of JSOn data is
 * Object and Array
 *
 * Example Object JSON
 *
 * {
 *		"FirstName": "Fajar",
 * 		"LastName": "Sidiq",
 *		"Age": 24
 * }
 */

/**
 * Object in Go-lang is represent in data type sturct
 * whereas every attribute in JSON Object is attribute in Struct
 *
 */

// Example Encode Struct to JSON
type SimpelCustomer struct {
	FirstName string
	LastName  string
	Age       int
}

func TestJSONSimpleObject(t *testing.T) {
	customer := SimpelCustomer{
		FirstName: "Fajar",
		LastName:  "Sidiq",
		Age:       24,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
