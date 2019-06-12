package main

import (
	"fmt"
	"strconv"
	"strings"
)

func strToIntArr(str string) (result []int){
	var k int
	arrStr := strings.Split(str,",")
	for _,each := range arrStr{
		k,_=strconv.Atoi(each)
		result = append(result,k)
	}
	return
}

func findMajority(arr []int) (result int,found bool) {
	limit := len(arr)/2
	count := make(map[int]int)
	for _,each := range arr{
		if count[each] > 0 {
			count[each]++
		} else {
				count[each] = 1
		}
		if count[each] > limit {
			return each,true
		}
	}
	return -1, false

}

func main()  {
	var numbers string
	fmt.Print("Input: ")
	_, _ = fmt.Scanln(&numbers)

	arrNum := strToIntArr(numbers)

	result, found := findMajority(arrNum)

	fmt.Print("Output: ")
	if !found{
		fmt.Println("Tidak Ditemukan")
	} else {
		fmt.Println(result)
	}


}

