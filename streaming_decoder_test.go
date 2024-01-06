package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

/** Streaming Decoder
 * previously we already study package json with conversion JSON data which
 * already in variable and string data or []byte
 *
 * in the fact, sometime JSON data comes from input in the form of io.Reader
 * (File, Network, Request Body)
 *
 * we can read all data first, then save in variable, and do conversion from JSON,
 * But there is actually no need to do that, because package json has feature to read
 * data from Stream
 */

func TestStreamDecoder(t *testing.T) {

	reader, _ := os.Open("Customer.json")
	decoder := json.NewDecoder(reader)

	customer := &Customer{}
	decoder.Decode(customer)

	fmt.Println(customer)
}
