package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countMedian(arr []int) float32 {
	if len(arr)%2 != 0 {
		return float32(arr[len(arr)/2])
	}
	return (float32(arr[len(arr)/2]) + float32(arr[(len(arr)/2)-1])) / 2
}

func main() {
	var numbers string
	fmt.Println("FIND MEDIAN OF AN ARRAY OF INTS")
	fmt.Print("Input (ex:1,2,3,4,5): ")
	_, _ = fmt.Scanln(&numbers)

	split := strings.Split(numbers, ",")
	var arrNum []int
	for _, each := range split {
		num, _ := strconv.Atoi(each)
		arrNum = append(arrNum, num)
	}

	fmt.Print("Output: ")
	fmt.Println(countMedian(arrNum))
}
