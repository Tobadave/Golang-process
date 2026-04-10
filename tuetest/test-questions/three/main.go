package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("ERROR: Required to input 2 strings")
		return
	}

	stringdata := os.Args[1]
	data := strings.Fields(stringdata)

	for _, w := range data {
		
	}


}
