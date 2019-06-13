package main

import (
	"fmt"
	"strconv"
	"strings"
)

func strToIntArr(str string) (result []int) {
	var k int
	arrStr := strings.Split(str, ",")
	for _, each := range arrStr {
		k, _ = strconv.Atoi(each)
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

func MergeSort(slice []int) []int {

	if len(slice) < 2 {
		return slice
	}
	mid := (len(slice)) / 2
	return Merge(MergeSort(slice[:mid]), MergeSort(slice[mid:]))
}

func HowManyGifts(maxBudget int, gifts []int) int {
	var sum int
	var count int
	for _, gift := range gifts {
		sum += gift
		if sum > maxBudget {
			break
		}
		count++
	}
	return count
}

func main() {
	var maxBudget int
	var gifts string
	fmt.Println("COUNT HOW MANY GIFTS CAN BE BOUGHT")
	fmt.Println("Input: ")
	fmt.Print("maxBudget = ")
	_, _ = fmt.Scanln(&maxBudget)
	fmt.Print("gifts = ")
	_, _ = fmt.Scanln(&gifts)

	sortedGifts := MergeSort(strToIntArr(gifts))

	fmt.Print("Output: ")
	fmt.Println(HowManyGifts(maxBudget, sortedGifts))
}
