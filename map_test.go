package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

/** Map
 * When using JSON, sometimes we found case dynamic JSON data
 * that is mean erratic, can increase, can decrease, and not static
 * In the case like that, use Struct will complicate, because in the Struct
 * we must ini specify all attribute
 *
 * for case like that, we can use data type map[string]interface{}
 * Automatically, "attribute" will be the key in the map, and "value" will be
 * value in the map
 *
 * Because the value is an interface{}, so we must do conversion manually
 * if we want to get a value
 *
 * Type data Map not support JSON tag again
 */
func TestMap(t *testing.T) {
	jsonRequest := `{"id":"ASDU2-12NC9AS-2321", "name": "samsung A 50s", "price": 3000000}`
	jsonBytes := []byte(jsonRequest)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "ASDU1G-JSNDFAD-2421I",
		"name":  "Apple Mac Book Pro M1",
		"price": "15000000",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
