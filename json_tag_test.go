package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

/** JSON Tag
 * Defautly attribute in the struct and JSON will mapping appropriate
 * with same attribute name (case sensitive)
 *
 * sometimes has style which different between attribute named
 * in the Struct and JSON, for example in the JSON we want to use snake_case,
 * but in the Struct, we want use PascalCase
 *
 * fortunately, package json support Tag Reflection
 * we want add tag reflection with json name, then following with attribute
 * which we want conversion from or to JSON
 */

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "id-0001",
		Name:     "ROG g 14",
		ImageUrl: "http://example.com",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id": "P0001", "name": "rog g-14", "image_url": "http://example.com"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}

	json.Unmarshal(jsonBytes, product)
	fmt.Println(product.Id)
	fmt.Println(product.Name)
	fmt.Println(product.ImageUrl)
}
