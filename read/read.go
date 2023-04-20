package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	var filename string
	name_slice := make([]name, 0)

	fmt.Println("File path:")
	fmt.Scan(&filename)

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		names := strings.Split(string(line), " ")

		name_slice = append(name_slice, name{fname: names[0], lname: names[1]})
	}

	fmt.Println("Read names")
	for _, names_struct := range name_slice {
		fmt.Printf("first name: %s; last name: %s\n", names_struct.fname, names_struct.lname)
	}
}
