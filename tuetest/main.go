package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//Read the files first
	data := Verify()
	// fmt.Println(data)

	if data == nil {
		fmt.Println("ERROR : Yout input does not align with the instructions")
		return
	}

	//the LOC below groups each ascii art into lines of 8, then new line (\n)
	lines := strings.Split(string(data), "\n")

	//puts the value of args.1 into input, as string
	input := string(os.Args[1]) //string in by default

	//this below sets the input as a string
	input = strings.ReplaceAll(input, "\\n", "\n") // str | old | new
	inputWords := strings.Split(input, "\n")

	// const height = 8 // the max height

	for _, word := range inputWords {
		if word == "" || word == "\n" {
			// fmt.Println("ERROR: You entered an empty string")
			fmt.Println() //should print a new line instead
			continue
		}

		for row := 1; row <= 8; row++ { //loops through the height

			for _, ch := range word {
				//char is now a rune since its each letter, in each word
				start := int((ch - 32) * 9)
				fmt.Print(lines[start+row])
			}

			fmt.Println() //prints a new line, after each row filling
		}
	}
}

func Verify() []byte { //fucntion that returns  string read file back

	// the length of the arguments must be 3 or 2
	if len(os.Args) < 2 || len(os.Args) > 3 { // Guards if the user input is not 3 or 2
		fmt.Println("File not up to length of 2 or 3")
		return nil
	}

	//selects the mode and optional
	readopt := ""

	//if the len is equal to 3
	if len(os.Args) == 3 { // include the output name...
		output := string(os.Args[2]) //the selected output

		if (output != "standard") && (output != "thinkertoy") && (output != "shadow") {
			fmt.Println("ERROR: Your file output option is Invalid!")
			return nil // dosent match the standard
		}

		readopt = output //pass the file
	} else {
		//default option
		readopt = "standard"
	}

	// the process of reading the file
	readfile := fmt.Sprintf("banner/%v.txt", readopt)

	read, err := os.ReadFile(readfile)
	if err != nil {
		fmt.Println("ERROR: One banner file not found")
		return nil
	}

	return read
}
