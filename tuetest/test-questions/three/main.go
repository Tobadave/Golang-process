package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("ERROR: Required to input 2 strings")
		return
	}

	// stringdata := os.Args[1]
	// data := strings.Fields(stringdata)

	assign, err := os.ReadFile("subs.txt")
	if err != nil {
		fmt.Println("ERROR: File not found!")
	}

	assign1 := string(assign)

	assign2 := strings.Split(assign1, "\n")
	fmt.Println(len(assign2))

	// for _, w := range data {

	// }
	// goes through the loop and
	//using regexp to find == word + = + word

	pattern := regexp.MustCompile(`s*\=s*`)
	match := pattern.FindAllSubmatch()
}
