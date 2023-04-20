package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string
	fmt.Printf("Type a string\n")
	fmt.Scan(&x)
	x = strings.ToLower(x)
	if strings.HasPrefix(x, "i") && strings.HasSuffix(x, "n") && strings.Contains(x, "a") {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not Found!\n")
	}
}
