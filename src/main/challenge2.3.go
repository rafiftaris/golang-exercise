package main

import (
	"fmt"
	"strings"
)

func main()  {
	var keys, values string
	fmt.Println("Input: ")
	fmt.Print("keys = ")
	_, _ = fmt.Scanln(&keys)
	fmt.Print("values = ")
	_, _ = fmt.Scanln(&values)

	keysArr := strings.Split(keys,",")
	valuesArr := strings.Split(values,",")

	var result = make(map[string]string)

	for i:=0 ; i<len(keysArr) ; i++{
		result[keysArr[i]] = valuesArr[i]
	}

	fmt.Println("Output: ",result)
}
