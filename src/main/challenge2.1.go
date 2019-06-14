package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func countMean(arr []int) int {
	var sum float64
	for _, each := range arr {
		sum += float64(each)
	}
	mean := sum / float64(len(arr))
	return int(math.Round(mean))
}

func main() {
	var numbers string
	fmt.Println("FIND MEAN OF AN ARRAY OF INTS")
	fmt.Print("Input (contoh:1,2,3,4,5): ")
	_, _ = fmt.Scanln(&numbers)

	split := strings.Split(numbers, ",")
	var arrNum []int
	for _, each := range split {
		num, _ := strconv.Atoi(each)
		arrNum = append(arrNum, num)
	}

	fmt.Print("Output: ")
	fmt.Println(countMean(arrNum))
}
