package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// fmt.Println("Working well")

	read, err := os.ReadFile("words.txt")
	if err != nil {
		fmt.Println("ERROR: File not found.")
		return
	}

	readString := string(read)
	trial := strings.Fields(readString)

	for i := 0; i < len(trial); i++ {
		ch := trial[i]
		num := i + 1
		// fmt.Println(fmt.Printf(`%v: %s`, num, ch))
		pre := fmt.Sprintf(`%v: %s`, num, ch)
		fmt.Println(string(pre))
		// fmt.Println(ch)
	}

}
