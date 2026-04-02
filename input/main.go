package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 { //checks if the argument is up to 2 and prints an error and ends the code
		fmt.Println("ERROR: Requires One input string!")
		return
	}

	filedata := verifyFile() // verifies and returns the "standard.txt" file in strings data
	if filedata == nil {     //guard error check
		fmt.Println("ERROR: Banner file not found")
		return
	}

	input := os.Args[1]                            //stores the value of strings as index of 1
	input = strings.ReplaceAll(input, "\\n", "\n") // this LOC replaces all literal "\n" with 'newline'
	inputText := strings.Split(input, "\n")        //gets the string and split all addup ifit sees a new line

	printAscii(inputText, filedata) //this function takes in our USER-INPUT and the BANNER "standard.txt" file
}

func verifyFile() []string {

	file, err := os.ReadFile("banner/standard.txt") //reads the txt file, in bytes by default
	if err != nil {                                 //guard error check
		return nil
	}
	//you dont need an else after a return (it already stops there)
	return strings.Split(string(file), "\n") //put all values of every line in an array
}

// this function checks if the
func printAscii(inputText, filedata []string) { //the USER-INPUT & FILEDATA are slices now

	const asciiOffset = 32 //the starting character of the ascii count
	const charHeight = 9   // the height of each character (row)

	for _, word := range inputText { //takes the input string at ONCE
		if word == "" { //checks of the whole string entered is just --> " "
			fmt.Println() //prints a new line
			continue
		}

		for row := 1; row <= 8; row++ { // 8 lines for each rows (height)

			for _, char := range word { //fills in the first row of every character before moving to the next line
				start := int(char-asciiOffset) * charHeight //uses 'char' to get the current rune to find the number on the banner file. * charHeight because this is not a regular single rune, one contains 9 line spaces, so 9 for each what ever the character is.
				fmt.Print(filedata[start+row])              // file data is the dict for our ascii characters of 9 lines, this fetches using 'start', and adds the value of 'row' (like adding +1, to go to the next line, so it dosent orint the sam row for all)
			}
			fmt.Println() // goes to the ext line after each line
		}
	}
}
