package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("ERROR: Not up to 2 argument")
		return
	}

	userinput := 

	read, _ := os.ReadFile("sample.txt")
	data := string(read)

	// fmt.Print((data))

	fmt.Println(len(data))

	sep := strings.Split(data, "\n") //takes by new line
	fmt.Println(len(sep))



	// order := strings.Fields(string(data)) //breaks by words,
}
