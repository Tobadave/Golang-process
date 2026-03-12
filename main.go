package main

import (
	// "errors"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: not enough args")
		os.Exit(1) //error 1 means incomplete return code, 0 means succesfull
	}

	content, err := os.ReadFile("Sample.txt")
	if err != nil {
		fmt.Println("ERROR 1 :", err)
		os.Exit(1)
	}

	fmt.Println("----------------------------")

	err = os.WriteFile("result.txt", []byte("You are going home"), 0644)
	if err != nil {
		fmt.Println("error writing file:", err)
		os.Exit(1)
	}

	newContent := string(content)
	fmt.Println(newContent)
	fmt.Println("-----Code succesful-------")
}
