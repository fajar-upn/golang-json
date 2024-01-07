package main

import (
	"encoding/json"
	"os"
	"testing"
)

/** Streaming Encoder
 * besides an decoder, package json support create encoder too which can use for
 * write JSON directly in the io.Writer
 *
 * Therefore, we don't need save JSON data firstly in the string variable or []byte,
 * we can directly write to io.Writer
 */

/** json.Encoder
 * for create Encoder, we can use function json.NewEncoder(<WRITER>)
 * and we write data as JSON directly to writer with function Encode(interfacer{})
 */

func TestStreamingEncode(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName: "Fajar",
		LastName:  "Sidiq",
		Age:       24,
		Married:   false,
	}

	encoder.Encode(customer)

}
