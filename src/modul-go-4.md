## Alta Go 3 # Struct & Method

### 🎯 Objektif

- Menerapkan Struct
- Menerapkan Method

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
  // input
}
```

### Challenge 2 - Two Sum

Given an array of integers, return indices of the two numbers such that they add up to a specific target. You may assume that each input would have exactly one solution, and you may not use the same element twice.

For example, if the array is [3, 5, 2, -4, 8, 11] and the sum is 7, your program should return [[11, -4], [2, 5]], because 11 + -4 = 7 and 2 + 5 = 7.

#### Sample Test Cases
```
Input: []int{3, 5, 2, -4, 8, 11}, 7
Output: [[11 -4] [2, 5]]
```

#### Code
```golang
package main

import(
  "fmt"
)

func main() {
  // kode disini
}
```

### Challenge 3 - Remove Set a Char From a String

The first thing to notice is that we’ll need to loop through the entire string S and loop through the array of characters because we’ll need to find the characters to actually remove from the string. There are two ways to construct the loops:

1. Loop through the array of characters and for each chracter, find all of the occurences in S and remove them.
2. Loop through the string S and if the current character exists in the array of characters somewhere, remove it.

We’re going to go with the second option because we can improve the algorithm by actually storing all of the characters in the array in a hash table. This will allow us to loop through the string S and determine if the current character needs to be removed by checking if it exists in the hash table. Remember from before, checking if an element exists in a hash table is done in constant time so the running time of our algorithm will be O(n) where n is the length of the string S plus some preprocessing where we loop through the array of characters and store them in a hash table.

#### Sample Test Cases
```
Input: []string{"h", "e", "w", "o"}, "hello world"
Output: "ll rld"
```

#### Code
```golang
package main

import(
  "fmt"
)

func main() {
  // kode disini
}
```

### Challenge 4 - Statistik

Tersedia 3 input:
1) kata (2 kemungkinan: 'min' atau 'max')
2) array berisi angka-angka positif
3) array berisi angka-angka positif

#### Sample Test Cases
```
Input: "max", []int{6, 2, 4, 10, 8, 2}, []int{4, 7, 9, 19}
Output: 10, 19

Input: "min", []int{5, 11, 18, 6}, []int{3, 9, 10, 13}
Output: 5 , 13
```

#### Code
```golang
package main

import(
  "fmt"
)

func main() {
  // kode disini
}
```

### Challenge 5 - How Many Gift

Toni menggambar sebuah segitiga yang alasnya berukuran A cm dan tingginya berukuran T cm. Ia ingin menghitung luas dari segitiga tersebut, tetapi ia lupa caranya, bantulah di Toni!

#### Sample Test Cases
```
Input: A = 20, T = 25
Output: 25
```

#### Code
```golang
package main

import(
  "fmt"
)

func main() {
  // input
  A := 20
  T := 25

  // kode disini
}
```

### Challenge 6 - Prima Segi Empat (BONUS)

Diberikan sebuah program yang menerima sebuah array angka. Program akan menghasilkan output modus dari array tersebut. Modus adalah angka dari sebuah deret yang paling banyak atau paling sering muncul. Contoh, modus dari [10, 4, 5, 2, 4] adalah 4.

Jika modus tidak ditemukan, program akan menghasilkan output -1, apabila ditemukan lebih dari dua nilai modus, tampilkan nilai modus yang paling pertama muncul (dihitung dari kiri ke kanan).

Dan apabila modus hanya memiliki 1 nilai yang sama maka program akan menghasilkan output -1, contohnya (1, 1, 1) adalah -1.

#### Sample Test Cases
```
Input: []int{10, 4, 5, 2, 4}
Output: 4

Input: []int{5, 10, 10, 6, 5}
Output: 5

Input: []int{1, 2, 3, 3, 4, 5}
Output: 3

Input: []int{10, 3, 1, 2, 5}
Output: -1

Input: []int{5, 5, 5, 5, 5}
Output: -1
```

#### Code
```golang
package main

import(
  "fmt"
)

func main() {
  // kode disini
}
```