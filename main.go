package main

import (
	// "errors"
	"bufio"
	"fmt"
	"os"
)

func main() {

	/**
	TRYING OUT PART I

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
	**/

	//TRYING OUT PART II
	file, err := os.Open("Sample.txt")
	if err != nil {
		fmt.Println("ERROR (OPEN Failed) : ", err)
		os.Exit(1) //bad launch
	}

	defer file.Close()

	scanner := bufio.NewScanner(file) //declaring a variable that stores in the reading value of 'file' above

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// package main
// import ("fmt"
// "strings")

// func main() {
//     // sentence := "it  was   the    best of times"
//     // fmt.Println(strings.Split(sentence, " "))
//     // fmt.Println(strings.Fields(sentence))

//     sentence := "it  was   the    best of times"
//     words := strings.Fields(sentence)
//     result := strings.Join(words, " ")
//     fmt.Println(result)
// }
