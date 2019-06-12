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

func Statistik(kata string, array1 []int, array2 []int) (int, int) {
	var max1, max2 int
	var min1 = 1<<(32 << (^uint(0) >> 32 & 1)-1) - 1
	var min2 = 1<<(32 << (^uint(0) >> 32 & 1)-1) - 1
	
	for _,each := range array1{
		if each > max1{
			max1 = each
		}
		if each < min1{
			min1 = each
		}
	}
	for _,each := range array2{
		if each > max2{
			max2 = each
		}
		if each < min2{
			min2 = each
		}
	}
	if kata == "max"{
		return max1,max2
	}
	if kata == "min"{
		return min1,min2
	}
	return 0, 0
}

func main()  {
	var keyword, nums1, nums2 string
	fmt.Println("Input: ")
	fmt.Print("kata = ")
	_,_ = fmt.Scanln(&keyword)
	fmt.Print("array1 = ")
	_,_ = fmt.Scanln(&nums1)
	fmt.Print("array2 = ")
	_,_ = fmt.Scanln(&nums2)

	arrNum1 := strToIntArr(nums1)
	arrNum2 := strToIntArr(nums2)

	fmt.Print("Output: ")
	fmt.Println(Statistik(keyword,arrNum1,arrNum2))
}

