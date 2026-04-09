# QUESTIONS & SAMPLE TESTS

**Here are 3 questions, each harder than the last** 🎯

## Question 1 — Warm up
Write a Go program that reads a file called words.txt, splits it by newlines, and prints each line with its line number in front.

Example

```bash
output:
1: Hello
2: World
3: Goodbye
```

## Question 2 — Think harder
Write a Go program that takes a word as a command line argument, reads a file called data.txt, and counts how many times that word appears in the file. Print the count at the end.
Example:
```go
go run . "hello"
//"hello" appears 3 times
```

## Question 3 — Challenge
Write a Go program that takes a sentence as a command line argument, and a filename as a second argument. It reads the file which contains a simple substitution map like:

```
hello = hi
world=earth
```

And replaces every matching word in the sentence with its substitution, then prints the result.
Example:
```
bashgo run . "hello world" "subs.txt"
// hi earth
```

Start with question 1 and paste your code when ready! Don't overthink it, you already know everything needed 🎯