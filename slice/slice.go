package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var typed_string string
	slice := make([]int, 3)
	for {
		fmt.Printf("Enter a number\n")
		if typed_string == "X" {
			break
		}
		fmt.Scan(&typed_string)
		if s, err := strconv.Atoi(typed_string); err == nil {
			slice = append(slice, s)
		}
	}
	sort.SliceStable(slice, func(i, j int) bool { return slice[i] < slice[j] })
	fmt.Println(slice)
}
