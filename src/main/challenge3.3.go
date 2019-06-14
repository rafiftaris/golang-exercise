package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringToArrayOfIntegers(str string) (result []int) {
	arrStr := strings.Split(str, ",")
	for _, each := range arrStr {
		k, err := strconv.Atoi(each)
		if err != nil {
			fmt.Println("Input Anda tidak sesuai dengan contoh. Harap ikuti contoh yang tertera.")
			os.Exit(2)
		}

		result = append(result, k)
	}
	return
}

func findMaxMin(keyword string, array []int) (result int){
	var min,max int
	max = array[0]
	min = array[0]

	for _, number := range array {
		if number > max {
			number = max
		}
		if number < min {
			number = min
		}
	}

	if keyword == "max"{
		return max
	}
	return min
}

func Statistik(kata string, array1 []int, array2 []int) (int, int) {
	resultFromArray1 := findMaxMin(kata,array1)
	resultFromArray2 := findMaxMin(kata,array2)
	return resultFromArray1,resultFromArray2
}

func main() {
	var keyword, numbersInputted1, numbersInputted2 string
	fmt.Println("FIND MAX/MIN OF TWO ARRAY OF INTS")
	for{
		fmt.Println("Input: ")
		fmt.Print("kata (max/min)= ")
		_, _ = fmt.Scanln(&keyword)
		fmt.Print("array1 (ex:1,2,3,4,5)= ")
		_, _ = fmt.Scanln(&numbersInputted1)
		fmt.Print("array2 (ex:1,2,3,4,5)= ")
		_, _ = fmt.Scanln(&numbersInputted2)

		if keyword != "max" && keyword != "min"{
			fmt.Println("kata harus merupakan 'max' atau 'min'")
			fmt.Println()
			continue
		}
		break
	}


	arrayOfIntegers1 := stringToArrayOfIntegers(numbersInputted1)
	arrayOfIntegers2 := stringToArrayOfIntegers(numbersInputted2)

	fmt.Print("Output: ")
	fmt.Println(Statistik(keyword, arrayOfIntegers1, arrayOfIntegers2))
}
