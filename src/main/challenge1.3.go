package main

import (
	"fmt"
)

func factor(num int) (result []int) {
	for i := 1; i <= num/2; i++ {
		if num%i == 0 {
			result = append(result, i)
		}
	}

	result = append(result, num)

	return
}

func printSlice(s []int) {
	for _, each := range s {
		fmt.Println(each)
	}
}

func main() {
	var num int
	fmt.Print("Input : ")
	_, _ = fmt.Scanln(&num)

	fmt.Println("Output :")
	printSlice(factor(num))
}
