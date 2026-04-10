package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: Required just a string")
		return
	}

	input := os.Args[1]

	data, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println("ERROR: File not found")
		return
	}

	data2 := string(data)
	readData := strings.Fields(data2)
	count := 0

	for _, word := range readData {
		if word == input {
			count++
		}
	}

	output := fmt.Sprintf(`%q is found %v times`, input, count)
	// fmt.Println(output)

	if count == 0 {
		fmt.Println("No match exist!")
	} else {
		fmt.Println(output)
	}
}
