package main

import (
	"fmt"
	"strings"
)

func main() {
	var keys, values string
	var keysArr, valuesArr []string
	fmt.Println("CONVERT CSV TO MAP")

	for {
		fmt.Println("Input: ")
		fmt.Print("keys (ex:name,age)= ")
		_, _ = fmt.Scanln(&keys)
		fmt.Print("values (ex:Aang,12)= ")
		_, _ = fmt.Scanln(&values)

		keysArr = strings.Split(keys, ",")
		valuesArr = strings.Split(values, ",")

		if keys == "" || values == ""{
			fmt.Println("Input tidak boleh kosong")
			fmt.Println()
			continue
		}
		if len(keysArr) != len(valuesArr){
			fmt.Println("Masukkan keys dan value dengan jumlah yang sama")
			fmt.Println()
			continue
		}

		break
	}

	var result = make(map[string]string)

	for i := 0; i < len(keysArr); i++ {
		result[keysArr[i]] = valuesArr[i]
	}

	fmt.Println("Output: ", result)
}
