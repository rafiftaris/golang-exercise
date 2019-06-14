package main

import (
	"fmt"
	"math"
)

func isPrime(num int) bool {
	for i := 2; i <= int(math.Ceil(float64(num)/2)); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// Generating prime numbers and put it inside the two dimensional array
func primaSegiEmpat(panjang, lebar, batasAwal int) {
	size := panjang * lebar
	rectangle := make([][]int, lebar)
	for i := 0; i < lebar; i++ {
		rectangle[i] = make([]int, panjang)
	}
	var index int
	var number = batasAwal + 1
	var sum int
	for index < size {
		for !isPrime(number) || number == 0 || number == 1{
			number++
		}
		rectangle[index/panjang][index%panjang] = number
		sum += number
		number++
		index++
	}
	printRectangle(rectangle)
	fmt.Println("Total: ", sum)
}

func printRectangle(rectangle [][]int)  {
	for i := 0; i < len(rectangle); i++ {
		for j := 0; j < len(rectangle[i]); j++ {
			fmt.Print(rectangle[i][j], " ")
		}
		fmt.Println()
	}
}

func main() {
	var panjang, lebar, batasAwal int
	fmt.Println("RECTANGULAR PRIME")
	for{
		fmt.Println("Input: ")
		fmt.Print("P = ")
		_, _ = fmt.Scanln(&panjang)

		if panjang <=0 {
			fmt.Println("P harus lebih besar dari 0")
			fmt.Println()
			continue
		}

		fmt.Print("L = ")
		_, _ = fmt.Scanln(&lebar)

		if lebar <=0 {
			fmt.Println("L harus lebih besar dari 0")
			fmt.Println()
			continue
		}

		fmt.Print("X = ")
		_, _ = fmt.Scanln(&batasAwal)

		if batasAwal < 0 {
			fmt.Println("X harus lebih besar atau sama dengan 0")
			fmt.Println()
			continue
		}
		break
	}


	primaSegiEmpat(panjang, lebar, batasAwal)
}
