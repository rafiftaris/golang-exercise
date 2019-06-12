package main

import (
	"fmt"
)

func isEven(num int) bool {
	return num%2==0
}

func main()  {
	var num int
	fmt.Print("Masukkan angka : ")
	_, _ = fmt.Scanln(&num)

	fmt.Print(num," adalah angka ")
	if isEven(num){
		fmt.Println("genap")
	} else {
		fmt.Println("ganjil")
	}
}

