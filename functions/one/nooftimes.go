package main

import "fmt"

func main() {
  fmt.Println("Try programiz.pro")
  
  sample := "hello..world"
//   fmt.Println(mytrim(sample))
  mytrim(sample)
}

func mytrim(s string) {
    look := map[rune]bool{
        '.': true,
        ',': true,
        ';': true,
    }

    result := map[rune]int{} // stores char and count

    for _, ch := range s {
        if _, ok := look[ch]; ok {
            result[ch]++ // increments count each time char is found
        }
    }
    
    for char, count := range result {
        fmt.Printf("%q: %d\n", char, count) //returns ----> ('.' : 2)
    }
}