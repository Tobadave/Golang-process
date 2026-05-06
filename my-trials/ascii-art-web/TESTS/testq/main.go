package main
import "fmt"
import "strings"
// import "sort"

func main() {
	sample := "ello how are you doing?"
	fmt.Println(wordsvcount(sample))

	one := "ehllo"
	fmt.Println(vowelstart(one))

	sample1 := "hello fam how"
	fmt.Println(reverseword(sample1))

	sample = "go is go and is fun"
	fmt.Println(remduplicate(sample))

	sample = "I love programming"
	fmt.Println(longest(sample))

	sample = "i am happy to be here"
	fmt.Println(max3(sample))

	// sample = "go is fun to learn"
	// fmt.Println(replacenw(sample, 2))

	// sample = "hello world"
	// fmt.Println(freq(sample))

	// sample := ""
	fmt.Println(anag("listen", "silent"))
	fmt.Println(anag("listen", "silent"))

	sample = "aaabbbccdd"
	fmt.Println(comp(sample))

	sample = "I use an umbrella"
	fmt.Println(emptyvowel(sample))

}
// Q9.
// Write a function that compresses repeated characters.
// "aaabbbccdd" → "a3b3c2d2"


// Q10.
// Write a function that takes a sentence and returns it with each word's vowels removed, except if removing all vowels would make the word empty.
// "I use an umbrella" → "I s n mbrll"

func emptyvowel(s string) []string {
	w := strings.Fields(s)
	build := []string{}

	for _, w := range w {
		newword := vowelremove(w)

		if len(newword) > 0 {
			build = append(build, newword)
		} else {
			continue
		}
	}

	return build
}



func comp(s string) string {
	store := map [rune]int{}
	var build strings.Builder
	// hold := []string{}

	for _, ch := range s {
		store[ch]++
	}

	for k, v := range store {
		one := ""
		try := string(k)
		one = fmt.Sprintf("%s%v", try, v)
		build.WriteString(one)

		// hold = append(hold, one)
		// sort(hold)
	}

	return build.String()
}








func anag(a, b string) bool {
	check := map[rune]bool{}

	if len(a) != len(b) {
		return false
	}

	for _, ch := range a{
		check[ch] = true
	}

	for _, ch := range b {
		if !check[ch]{
			return false
		}
	}
	return true
}
// func freq(s string) int {
// 	store := map[rune]int{}
// 	var res int

// 	for _, ch := range s {
// 		store[ch]++ //counts in map
// 	}

// 	for k, v := range store {
// 		const max := v

// 		if store[k]> max {
// 			max = store[k]
// 		}

// 		res = v
// 	}

// 	// res = k
// 	return res
// }

// Q6.
// Write a function that replaces every nth word with "*".
// everynth("go is fun to learn", 2) → "go * fun * learn"



// Q3.
// Write a function that removes all duplicate words from a sentence (keep first occurrence).
// "go is go and is fun" → "go is and fun"


// Q4.
// Write a function that returns the longest word in a string.
// "I love programming" → "programming"


func max3(s string) string {
	w := strings.Fields(s)

	for i, ch := range w {
		if len(ch) > 3 {
			w[i] = strings.Title(ch)
		}
	}

	return strings.Join(w, " ")
}


func longest(s string) string {
	w := strings.Fields(s)
	max := w[0] //string

	for _, w := range w {
		if len(w) > len(max) {
			max = w
		}
	}

	return max
}

func remduplicate(s string) []string {
	words := strings.Fields(s)
	checkbox := []string{}
	try := map[string]int{}

	for _, w := range words {
		try[w]++
	}

	fmt.Println(try)

	for _, w := range words {
		if (try[w] > 1) && (testrun(w, checkbox)){
		// if (try[w] > 1) && slices.Contains(w, checkbox){

			continue
		} else if try[w] >= 1 {
			checkbox = append(checkbox, w)
		} 
	}

	return checkbox
}

func testrun(s string, full []string) bool {
	full1 := strings.Join(full, " ")

	w := strings.Fields(full1)
	for _, word := range w {
		if word == s {
			return true
		}
	}

	return false


	//all the above can be re-written as 
	// : return slices.Contains(words, s) --> true or false 
}


// func testrun(s string, full []string) bool {
// 	full1 := strings.Join(full, " ")

// 	w := strings.Fields(full1)
// 	return slices.Contains(w, s)
// }


// func remduplicate(s string) []string {
// 	words := strings.Fields(s)
// 	checkbox := []string{}
// 	try := map[string]bool{}

// 	for _, w := range words {
// 		try[w] = true
// 	}

// 	fmt.Println(try)

// 	for _, w := range words {
// 		// if try[w] {
// 		// 	continue
// 		// } else {
// 		// 	checkbox = append(checkbox, w)
// 		// }

// 		if _, ok := try[w]; ok {
// 			checkbox = append(checkbox, w)
// 		} else {
// 			continue
// 		}
// 	}

// 	return checkbox
// }






func reverseword(s string) []string {
	words := strings.Fields(s)
	revwords := []string{}

	for _, w := range words { //takes "hello"

		new := ""

		for i:=len(w)-1; i>=0; i-- { //adds from the back "olleh"
			new += string(w[i])
		}
		revwords = append(revwords, new)
	}

	return revwords
}



func wordsvcount(s string) int {
	words := strings.Fields(s)
	count := 0

	for _, w := range words {
		if vowelstart(w) {
			count++
		}
	}

	return count
}


func vowelstart(s string) bool {
	vowels := map [rune]bool {
		'a':true,
		'e':true,
		'i':true,
		'o':true,
		'u':true,
	}

	first := rune(s[0])
	if _, ok := vowels[first]; ok {
		return true
	}

	return false
}

func vowelremove(s string) string {
	var build strings.Builder

	vowels := map [rune]bool {
		'a':true,
		'e':true,
		'i':true,
		'o':true,
		'u':true,
	}

	first := rune(s[0])
	if _, ok := vowels[first]; ok {
		f := string(first)
		build .WriteString(f)
	}

	ans := build.String()

	return ans
}

