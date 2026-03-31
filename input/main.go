package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	input, err := os.ReadFile("banner/standard.txt")
	if err != nil {
		fmt.Println("Not working")
		return
	}

	lines := strings.Split(string(input), "/n")
	fmt.Println(lines)
}
