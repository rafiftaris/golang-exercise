package main

import (
	"fmt"
)

func isPalindrome(word string) bool {
	runes := []rune(word)
	j := len(runes)-1
	for i:=0 ; i<j ; i++{
		if runes[i] != runes[j]{
			return false
		}
		j--
	}
	return true
}

func main(){
	var word string
	fmt.Print("Input: ")
	_, _ = fmt.Scanln(&word)

	fmt.Print("Output: ")

	if !isPalindrome(word){
		fmt.Print("Bukan ")
	}
	fmt.Println("Palindrom")
}

