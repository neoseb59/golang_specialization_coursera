package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	person_map := make(map[string]string)
	var typed_string string

	fmt.Println("Enter your name")
	fmt.Scan(&typed_string)
	person_map["name"] = typed_string

	fmt.Println("Enter your address")
	fmt.Scan(&typed_string)
	person_map["address"] = typed_string

	json_map, _ := json.Marshal(person_map)
	fmt.Println(string(json_map))
}
