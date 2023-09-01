package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

/**
 * Go-lang already has function for do conversion data to json, that is
 * use function json.Marshall(interface{})
 *
 * Because parameter is a interface{}, so we can insert any data type into
 * Marshal function
 */

func logJSON(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestEncode(t *testing.T) {
	logJSON("Fajar")
	logJSON(1)
	logJSON(true)
	logJSON([]string{"Muhammad", "Fajar"})
}
