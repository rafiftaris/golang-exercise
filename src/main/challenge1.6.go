package main

import (
	"fmt"
)

func asterixTriangle(k int)  {
	for i:=1 ; i<=k ; i++{
		for j:=0 ; j<k-i ; j++{
			fmt.Print(" ")
		}
		for n:=1 ; n<i ; n++{
			fmt.Print("*")
			if n<i-1{
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main()  {
	var k int
	fmt.Print("Input: ")
	_, _ = fmt.Scanln(&k)

	fmt.Print("Output: ")
	asterixTriangle(k)
}

