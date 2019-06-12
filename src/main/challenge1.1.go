package main

import (
	"fmt"
)

func luasSegitiga(alas float32, tinggi float32) float32{
	return 0.5*alas*tinggi
}

func main()  {
	var input1, input2 float32
	fmt.Print("Masukkan alas dan tinggi segitiga : ")
	_, _ = fmt.Scanln(&input1, &input2)

	fmt.Println("Luas segitiga : ",luasSegitiga(input1,input2),"cm")

}

