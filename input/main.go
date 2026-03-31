package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	read, err := os.ReadFile("banner/standard.txt")
	if err != nil {
		fmt.Println("Not working")
		return
	}

	lines := strings.Split(string(read), "\n")
	// fmt.Println(lines)

	if len(os.Args) != 2 {
		fmt.Println("Required just 2 Arguments for this function!")
		return
	}

	inputText := strings.Split(os.Args[1], "\n")

	for _, word := range inputText {
		if word == "" { //if empty
			fmt.Println()
			continue
		}

		for row := 1; row <= 8; row++ {
			for _, char := range word { //gets each letter in the word
				start := int(char-32) * 9 //get the line proximity
				fmt.Print(lines[start+row])
			}
			fmt.Println()
		}
	}
}
