package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 && len(os.Args) != 3 { //checks if the argument is up to 2 and prints an error and ends the code
		fmt.Println("ERROR: Requires One OR Two input string!")
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

	fmt.Println(printAscii(inputText, filedata)) //this function takes in our USER-INPUT and the BANNER "standard.txt" file
}

func verifyFile() []string { // checks and vetifies the banner files and user inout read

	fileopt := "" //a new var to store value of input

	if len(os.Args) == 3 { //if the user inputs 3 characters
		fileopt = os.Args[2] // checks for the value of the 3rd index
	} else {
		fileopt = "standard" //default file name if user dosent fill 3rd place
	}

	if (fileopt != "standard") && (fileopt != "thinkertoy") && (fileopt != "shadow") { //must be either of them
		fmt.Print("ERROR : Invalid Banner name & ") // if none of the names is entered by the user
		return nil                                  // returns nil
	}

	// fileopt = fmt.Sprintf(`banner/%v.txt`, fileopt)

	file, err := os.ReadFile(fmt.Sprintf(`banner/%v.txt`, fileopt)) //reads the txt file, in bytes by default
	if err != nil {                                                 //guard error check
		return nil
	}
	//you dont need an else after a return (it already stops there)
	return strings.Split(string(file), "\n") //put all values of every line in an array
}

// this function checks if the
func printAscii(inputText, filedata []string) string { //the USER-INPUT & FILEDATA are slices now
	var sb strings.Builder

	const asciiOffset = 32 //the starting character of the ascii count
	const charHeight = 9   // the height of each character (row)

	for _, word := range inputText { //takes the input string at ONCE
		if word == "" { //checks of the whole string entered is just --> " "
			sb.WriteString("\n") //prints a new line
			continue
		}

		for row := 1; row <= 8; row++ { // 8 lines for each rows (height)

			for _, char := range word { //fills in the first row of every character before moving to the next line
				start := int(char-asciiOffset) * charHeight //uses 'char' to get the current rune to find the number on the banner file. * charHeight because this is not a regular single rune, one contains 9 line spaces, so 9 for each what ever the character is.
				sb.WriteString(filedata[start+row])         // file data is the dict for our ascii characters of 9 lines, this fetches using 'start', and adds the value of 'row' (like adding +1, to go to the next line, so it dosent orint the sam row for all)
			}
			sb.WriteString("\n") // goes to the ext line after each line
		}
	}

	return sb.String()
}
