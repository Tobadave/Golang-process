package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Working well")

	read, err := os.ReadFile("words.txt")
	if err != nil {
		fmt.Println("ERROR: File not found.")
		return
	}

	readString := string(read)

	fmt.Println(readString)
	read2 := strings.Fields(readString)

	fmt.Println(strings.Join(read2, " "))

	//function to test it

	// for i:=0; i<len(readString); i++ {
	// 	ch := string(readString[i])
	// 	out := strings.TrimSpace(ch)

	// 	fmt.Println(out)
	// }

	// fmt.Println(out)
}
