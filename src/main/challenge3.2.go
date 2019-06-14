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

func TwoSums(Nums []int, Sum int) (result [][]int) {
	for i := 0; i < len(Nums); i++ {
		for j := i + 1; j < len(Nums); j++ {
			if Nums[i]+Nums[j] == Sum {
				result = append(result, []int{Nums[i], Nums[j]})
			}
		}
	}
	return
}

func main() {
	var nums string
	var sum int
	fmt.Println("FIND TWO INTS WHICH SUMS EQUALS TO 'SUM'")
	fmt.Println("Input: ")
	fmt.Print("Nums (ex:1,2,3,4,5)= ")
	_, _ = fmt.Scanln(&nums)
	fmt.Print("Sum = ")
	_, _ = fmt.Scanln(&sum)

	arrNums := strToIntArr(nums)

	fmt.Print("Output: ")
	fmt.Println(TwoSums(arrNums, sum))

}
