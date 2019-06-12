package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countMean(arr []int) int {
	var sum int
	for _,each:=range arr{
		sum += each
	}
	return sum/len(arr)
}

func main()  {
	var numbers string
	fmt.Print("Input: ")
	_, _ = fmt.Scanln(&numbers)

	split := strings.Split(numbers,",")
	var arrNum []int
	for _,each := range split{
		num,_:=strconv.Atoi(each)
		arrNum = append(arrNum,num)
	}

	fmt.Print("Output: ")
	fmt.Println(countMean(arrNum))
}

