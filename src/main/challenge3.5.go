package main

import (
	"fmt"
	"math"
)

func checkPrime(num int) bool {
	for i := 2; i <= int(math.Ceil(float64(num)/2)); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primaSegiEmpat(P, L, X int) {
	size := P * L
	arr := make([][]int, L)
	for i := 0; i < L; i++ {
		arr[i] = make([]int, P)
	}
	var count int
	var num = X + 1
	var sum int
	for count < size {
		for !checkPrime(num) || num == 0 || num == 1{
			num++
		}
		arr[count/P][count%P] = num
		sum += num
		num++
		count++
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println("Total: ", sum)
}

func main() {
	var p, l, x int
	fmt.Println("RECTANGULAR PRIME")
	for{
		fmt.Println("Input: ")
		fmt.Print("P = ")
		_, _ = fmt.Scanln(&p)

		if p <=0 {
			fmt.Println("P harus lebih besar dari 0")
			fmt.Println()
			continue
		}

		fmt.Print("L = ")
		_, _ = fmt.Scanln(&l)

		if l <=0 {
			fmt.Println("L harus lebih besar dari 0")
			fmt.Println()
			continue
		}

		fmt.Print("X = ")
		_, _ = fmt.Scanln(&x)

		if x < 0 {
			fmt.Println("X harus lebih besar atau sama dengan 0")
			fmt.Println()
			continue
		}
		break
	}


	primaSegiEmpat(p, l, x)
}
