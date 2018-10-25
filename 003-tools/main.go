package main

import "fmt"
import "github.com/jwalkerdev/learning-go-class/003-tools/mypkg"

func main() {
	word1 := "level"
	fmt.Printf("palindrome test: %s -> %v\n", word1, mypkg.IsPalindrome(word1))
}
