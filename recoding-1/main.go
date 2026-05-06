package main

import (
	"fmt"
	// "strconv"
	// "strings"
)

func main() {
	// one := CeaserCipher("Hello", 3)
	// fmt.Println(one)
}

// Determine if a string's brackets — (), [], {} — are properly balanced and correctly nested.

// isBalanced("({[]})") → true isBalanced("([)]") → false

func brackets(s string) bool {
	brack1 := map[rune]bool {
		'[':true,
		'{':true,
		'(':true,
	}

	
	brack2 := map[rune]bool {
		']':true,
		'}':true,
		')':true,
	}

	all := map[rune]bool {
		']':true,
		'}':true,
		')':true,
		'[':true,
		'{':true,
		'(':true,
	}


	// for _, ch := range s {
	// 	if val, ok := brack[ch]; ok {
	// 		if ( ch != val) && 

	// 	}
	// }
	build := []rune{}
	for _, ch := range s {
		if _, ok := all[ch]; ok {
			build = append(build, all[ch])
		}
	}

	brac := map[rune]int {}
	// newone := br
	for _, ch := range s {
		if _, ok  := brack1[ch]; ok { //found one bracket...
			brac[ch]++
		}

		if _, ok  := brack2[ch]; ok { //found one bracket...

			brac[ch]++ //they should be equal to them selvues
		}


	}
}

// func CeaserCipher(s string, n int) string {
// 	// var build strings.Builder
// 	build := ""
// 	total := 26
// 	smallmin, smallmax := 65, 90
// 	uppermin, uppermax := 110, 135

// 	for _, ch := range s {
// 		curr := rune(ch)

// 		if curr >= smallmin && curr <= smallmax {
// 			new := strconv.Itoa(smallmax + ((curr-smallmin)%total))
// 			build += new

// 			fmt.Println(build)
// 			// fmt.Println(build.String())
// 		} else if curr >= uppermin && curr <= uppermin {
// 			new := strconv.Itoa(uppermax + ((curr-uppermin)%total))
// 			build += new
// 		} else {
// 			// build.WriteString(strconv.Itoa(curr))
// 			build += strconv.Itoa(curr)
// 		}
// 	}

	// res := build.String()

	return build
}

// 65 - 90 --> small letters
// 110 - 122 --? upper case
