package main

import (
	"fmt"
	"math"
)

func checkPrime(num int) bool{
	for i:=2 ; i<int(math.Ceil(float64(num)/2)) ; i++{
		if num%i==0{
			return false
		}
	}
	return true
}

func main()  {
	var num int
	fmt.Print("Input: ")
	_, _ = fmt.Scanln(&num)

	fmt.Print("Output: ")
	if !checkPrime(num){
		fmt.Print("Bukan ")
	}
	fmt.Println("Bilangan Prima")
}

