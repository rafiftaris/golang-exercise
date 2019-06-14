package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func removeChar(str string, chars []string) (result string) {
	var exist bool
	runes := []rune(str)
	for _, each := range runes {
		exist = false
		for _, char := range chars {
			if each == []rune(char)[0] {
				exist = true
				break
			}
		}
		if !exist {
			result += string(each)
		}
	}
	return
}

func main() {
	var str, chars string
	in := bufio.NewReader(os.Stdin)
	fmt.Println("REMOVE CHARACTERS FROM A STRING")
	fmt.Println("Input: ")
	fmt.Print("String = ")
	str, _ = in.ReadString('\n')
	fmt.Print("Char (ex:a,b,c,d)= ")
	chars, _ = in.ReadString('\n')

	arrChar := strings.Split(chars, ",")

	fmt.Println(removeChar(str, arrChar))
}
