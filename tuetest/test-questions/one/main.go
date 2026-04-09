// Write a Go program that reads a file called words.txt, splits it by newlines, and prints each line with its line number in front.
// Example output:
// 1: Hello
// 2: World
// 3: Goodbye

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("words.txt")
	if err != nil {
		fmt.Println("ERROR: File not found")
		return
	}

	data := string(file)

	new := strings.Split(data, " ")

	// fmt.Println(data)
	// fmt.Println(new)
	// fmt.Println(len(new))

	for i := 0; i < (len(new)); i++ {

		text := (new[i])
		pre := i + 1
		output := fmt.Sprintf(`%v : %s`, pre, text)
		fmt.Println(output)
	}
}
