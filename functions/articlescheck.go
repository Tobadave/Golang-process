package articlescheck

import (
	"strings"
)

func punc(s string) string {
	sects := strings.Fields(s) // breaks "pick", "an", "orange" -> 0, 1, 2
	list := []string{}         //an empty string for new list

	for i := 0; i < len(sects); i++ {
		curr := string(sects[i]) //current char = pick
		prev := ""               //decling prev. isInvalid on first index

		if len(list) > 0 {
			prev = string(sects[i-1])
		}
		//invalid

		if (len(list) != 0) && checkvow(string(curr[0])) && (prev == "a") {
			list[len(list)-1] = "an"
			list = append(list, curr)
		} else {
			list = append(list, curr)
		}
	}

	output := strings.Join(list, " ")
	return output
}

func checkvow(s string) bool {
	vowels := map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
	}

	if _, ok := vowels[s]; ok {
		return true
	} else {
		return false
	}
}
