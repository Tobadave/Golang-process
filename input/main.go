package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 { //checks if the argument is up to 2
		fmt.Println("ERROR: Requires One input string!")
		return
	}

	file, err := os.ReadFile("banner/standard.txt") //reads the txt file, in bytes by default
	if err != nil {                                 //guard error check
		fmt.Println("ERROR: Banner file not found")
		return
	}

	filedata := strings.Split(string(file), "\n") //put all values of every line in an array
	input := strings.Split(os.Args[1], "\n")      //gets the string

	for _, word := range input {
		if word == "" {
			fmt.Println()
			continue
		}

		for row := 1; row <= 8; row++ {

			for _, ch := range word {
				start := int(ch-32) * 9
				fmt.Print(filedata[start+row])
			}
			fmt.Println()
		}
	}

}
