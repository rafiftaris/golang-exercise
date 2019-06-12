## Modul 3 # Function

### 🏠 [back to home](https://github.com/alterra-academy/golang-class)

### 🎯 Objektif

- Menerapkan Function untuk menyelesaikan masalah pemrograman

### ✏️ Challenge

### Challenge 1 - Bilangan Prima

Dalam matematika, bilangan prima adalah bilangan asli yang lebih besar dari angka 1, yang faktor pembaginya adalah 1 dan bilangan itu sendiri. 2 dan 3 adalah bilangan prima. 4 bukan bilangan prima karena 4 bisa dibagi 2. Kamu diminta untuk membuat fungsi untuk menentukan bahwa sebuah bilangan termasuk bilangan prima atau tidak.

#### Sample Test Cases
```
Input: 3
Output: Bilangan Prima

Input: 5
Output: Bilangan Prima

Input: 6
Output: Bukan Bilangan Prima

Input: 7
Output: Bilangan Prima
```

#### Code
```golang
package main

import(
  "fmt"
)

func Prima(N int) string {

}

func main() {
  fmt.Println(Prima(3)) // Bilangan Prima
  fmt.Println(Prima(5)) // Bilangan Prima
  fmt.Println(Prima(6)) // Bukan Bilangan Prima
  fmt.Println(Prima(7)) // Bilangan Prima
}
```

### Challenge 2 - Two Sum

Anda akan diberikan sebuah Array bilangan integer (Nums), tugas anda adalah menemukan pasangan nilai yang jika dijumlah akan menghasilkan nilai Sum. Pada kasus ini anda tidak dapat menggunakan elemen yang sama untuk dua kali.

Misalnya, jika Nums adalah [3, 5, 2, -4, 8, 11] dan (Sum) jumlahnya adalah 7, program Anda harus mengembalikan [[11, -4], [2, 5]], karena 11 + - 4 = 7 dan 2 + 5 = 7.

#### Sample Test Cases
```
Input: Nums = []int{2, 7, 11, 15}, Sum = 9
Output: [[2 9]]

Input: Nums = []int{3, 5, 2, -4, 8, 11}, Sum = 7
Output: [[11 -4] [2, 5]]

Input: Nums = []int{1, 2, 4, 6, 10, 5, 13, 8, 14, 5}, Sum = 10
Output: [[2 8] [4 6] [5 5]]
```

#### Code
```golang
package main

import(
  "fmt"
)

func TwoSum(Nums []int, Sum int) [][]int {

}

func main() {
  // kode disini
}
```

### Challenge 3 - Statistik

Program statistik adalah program yang menerima 3 input:
1. kata (2 kemungkinan: 'min' atau 'max')
2. array 1 berisi angka-angka positif
3. array 2 berisi angka-angka positif

Buatlah sebuah program untuk menampilkan nilai maksimal apabila kata = "max" atau minimal apabila kata = "min" dari inputan dua array 1 dan 2!

#### Sample Test Cases
```
Input: kata = "max", array1 = []int{6, 2, 4, 10, 8, 2}, array2 = []int{4, 7, 9, 19}
Output: 10, 19

Input: kata = "min", array1 = []int{5, 11, 18, 6}, array2 = []int{3, 9, 10, 13}
Output: 5, 13
```

#### Code
```golang
package main

import(
  "fmt"
)

func Statistik(kata string, array1 []int, array2 []int) (int, int) {

}

func main() {
  // kode disini
}
```

### Challenge 4 - How Many Gift

Andi meminta Rana untuk membeli daftar oleh-oleh saat trip berikutnya, sekarang Rana ingin tahu berapa jumlah oleh-oleh paling banyak yang bisa dia beli. Implementasikan function dibawah untuk membantu Rana!

Pada function HowManyGifts(maxBudget, gifts). Parameter pertama adalah budget Rana, yang kedua adalah sebuah Array yang berisi harga setiap oleh-oleh. Function ini harus mengembalikan nilai yang mewakili jumlah maksimum oleh-oleh yang Rana dapat beli.


#### Sample Test Cases
```
Input: maxBudget = 25000, gifts = []int{20000, 5000, 10000, 6000, 4000}
Output: 4
```

Output adalah 4, karena bisa membeli oleh-oleh dengan harga 5000, 10000, 6000, 4000

#### Code
```golang
package main

import(
  "fmt"
)

func HowManyGifts(maxBudget, gifts) int {

}

func main() {
  fmt.Println(HowManyGifts(30000, []int{15000, 12000, 5000, 3000, 10000})) // 4
  fmt.Println(HowManyGifts(10000, []int{2000, 2000, 3000, 1000, 2000, 10000})) // 5
  fmt.Println(HowManyGifts(4000, []int{7500, 1500, 2000, 3000})) // 2
  fmt.Println(HowManyGifts(50000, []int{25000, 25000, 10000, 15000})) // 3
  fmt.Println(HowManyGifts(0, []int{10000, 3000})) // 0
}
```

### Challenge 5 - Prima Segi Empat

Buatlah segiempat berukuran P x L yang berisikan bilangan bilangan prima setelah X, pada bagian akhir jumlahkan seluruh bilangan prima tersebut. Anda diperbolehkan menggunakan lebih dari satu function untuk menyelesaikan masalah ini!

#### Sample Test Cases
```
Input: P = 2, L = 3, X = 13
Output:
17 19
23 29
31 37
Total: 156

Input: P = 5, L = 2, X = 1
2  3  5  7 11
13 17 19 23 29
Total: 129
```

#### Code
```golang
package main

import(
  "fmt"
)

func primaSegiEmpat(P, L, X int) {

}

func main() {
  // Driver Code
  primaSegiEmpat(2, 3, 13)
  /*
  17 19
  23 29
  31 37
  Total: 156
  */

  primaSegiEmpat(5, 2, 1)
  /*
  2  3  5  7 11
  13 17 19 23 29
  Total: 129
  */
}
```