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

func findModus(arr []int) (result int) {
	var max = 0
	dict := make(map[int]int)
	for _, each := range arr {
		if dict[each] > 0 {
			dict[each]++
		} else {
			dict[each] = 1
		}
		if dict[each] > max {
			max = dict[each]
		}
	}

	found := false
	for key, value := range dict {
		if max == value {
			if !found {
				found = true
				result = key
			}
		}
	}
	return
}

func main() {
	var numbers string
	fmt.Println("FIND MODUS OF AN ARRAY OF INTS")
	fmt.Print("Input (ex:1,2,3,4,5): ")
	_, _ = fmt.Scanln(&numbers)

	arrNum := strToIntArr(numbers)

	fmt.Print("Output: ")
	fmt.Println(findModus(arrNum))

}
