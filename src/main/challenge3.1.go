package main

import (
	"fmt"
	"math"
)

func Prima(N int) string{
	for i:=2 ; i<int(math.Ceil(float64(N)/2)) ; i++{
		if N%i==0{
			return "Bukan Bilangan Prima"
		}
	}
	return "Bilangan Prima"
}

func main()  {
	var num int
	fmt.Print("Input: ")
	_, _ = fmt.Scanln(&num)

	fmt.Print("Output: ")
	fmt.Println(Prima(num))
}
