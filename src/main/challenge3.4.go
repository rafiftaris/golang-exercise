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

// Merges left and right slice into newly created slice
func Merge(left, right []int) []int {

	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}

// Breaks down slice into two smaller pieces (recursively) then merge sort the slices
func MergeSort(slice []int) []int {

	if len(slice) < 2 {
		return slice
	}
	mid := (len(slice)) / 2
	return Merge(MergeSort(slice[:mid]), MergeSort(slice[mid:]))
}

func HowManyGifts(maxBudget int, gifts []int) int {
	var priceSum int
	var giftBought int
	for _, giftPrice := range gifts {
		priceSum += giftPrice
		if priceSum > maxBudget {
			break
		}
		giftBought++
	}
	return giftBought
}

func main() {
	var maxBudget int
	var gifts string
	fmt.Println("COUNT HOW MANY GIFTS CAN BE BOUGHT")
	for {
		fmt.Println("Input: ")
		fmt.Print("maxBudget = ")
		_, _ = fmt.Scanln(&maxBudget)
		fmt.Print("gifts (ex:10000,20000,40000)= ")
		_, _ = fmt.Scanln(&gifts)

		if maxBudget < 0{
			fmt.Println("maxBudget harus lebih besar atau sama dengan 0")
			fmt.Println()
			continue
		}
		break
	}

	arrayOfGifts := stringToArrayOfIntegers(gifts)
	sortedGifts := MergeSort(arrayOfGifts)

	fmt.Print("Output: ")
	fmt.Println(gifts)
	fmt.Println(HowManyGifts(maxBudget, sortedGifts))
}
