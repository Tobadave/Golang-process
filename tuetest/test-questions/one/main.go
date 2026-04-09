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
	new2 := strings.Join(new, " ")
	tn := strings.Trim(new2, " ")

	fmt.Println(new)
	fmt.Println(new2)
	fmt.Println(tn)
	fmt.Println("=======")
	fmt.Println(len(new))
	fmt.Println(len(new2))
	fmt.Println((len(tn)))

	// fmt.Println(data)
	// fmt.Println(new)
	// fmt.Println(len(new))

	// works but does not guard
	//-----------------------
	// for i := 0; i < (len(new)); i++ {

	// 	text := (new[i])
	// 	pre := i + 1
	// 	output := fmt.Sprintf(`%v : %s`, pre, text)
	// 	fmt.Println(output)
	// }
	fmt.Println(printNo(new2))
	// printNo(new2)
}

func printNo(s string) string {
	final := ""

	for i := 0; i < (len(s)); i++ {

		text := (s[i])
		pre := i + 1
		output := fmt.Sprintf(`%v : %s`, pre, text) 
		// fmt.Println(output)
		// return output
		final += string(output)
	}
	return final
}
