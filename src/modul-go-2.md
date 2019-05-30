## Modul 2 # Array, Slice & Map

### 🏠 [back to home](https://github.com/alterra-academy/golang-class)

### 🎯 Objektif

- Menerapkan Array
- Menerapkan Slice
- Menerapkan Map

### ✏️ Challenge

### Challenge 1 - Mean

Diberikan sebuah program yang menerima sebuah array angka. Program akan menampilkan mean dari array angka tersebut.
Mean adalah angka rata-rata dari deret bilangan tersebut. Contoh, mean dari [1, 2, 3, 4, 5] adalah 3. Kita perlu kemudian melakukan pembulatan angka dari hasil mean yang didapatkan.

#### Sample Test Cases
```
Input: []int{1, 2, 3, 4, 5}
Output: 3

Input: []int{3, 5, 7, 5, 3}
Output: 5

Input: []int{6, 5, 4, 7, 3}
Output: 5
```

#### Code
```golang
package main

import(
  "fmt"
)

func main() {
  // input
}
```

### Challenge 2 - Median

Diberikan sebuah array bertipe angka, program akan mencetak median dari deret angka tersebut. Median adalah nilai tengah dari sebuah deret bilangan. Contoh, median atau dari [1, 2, 3, 4, 5] adalah 3 yang merupakan nilai yang ada di posisi tengah dari deret tersebut.

Median dari deret yang berjumlah genap adalah rata-rata dari dua nilai tengah. Contoh, median dari [1, 2, 3, 4] adalah 2.5, karena (2 + 3 / 2). Pada kasus ini nilai dari angka tidak perlu diurutkan.

Catatan: pada kasus ini kamu tidak perlu melakukan sorting terhadap elemen seperti konsep matematika.


#### Sample Test Cases
```
Input: []int{1, 2, 3, 4, 5}
Output: 3

Input: []int{1, 3, 4, 10, 12, 13}
Output: 7

Input: []int{7, 7, 8, 8}
Output: 7.5
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

### Challenge 3 - CSV to Map

Buatlah sebuah merubah string yang diberikan menjadi sebuah Map baru. Program menerima 2 input yaitu keys dan values yang merupakan sebuah CSV (Comma Separated Values) Keys & values di dalam Map baru nantinya akan diambil dari kedua input yang diberikan.

#### Sample Test Cases
```
Input: keys = "name,age", values = "Aang,112"
Output: map[name:Dimitri phoneNumber:+666123654]

Input: keys = "firstName,lastName,nationality", values = "Sergei,Dragunov,Russia"
Output: map[firstName:Sergei lastName:Dragunov nationality:Russia]
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

### Challenge 4 - Remove Set of Character From a String

Kamu akan diiberi sebuah String dan Array of Char. Buatlah sebuah program untuk menghapus Char dalam String, Char yang perlu dihapus adalah Char yang muncul dalam Array of Char.

#### Sample Test Cases
```
Input: String = "hello world", Char = []byte{'h', 'e', 'w', 'o'}
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

### Challenge 5 - Modus

Diberikan sebuah program yang menerima sebuah array angka. Program akan menghasilkan output modus dari array tersebut. Modus adalah angka dari sebuah deret yang paling banyak atau paling sering muncul. Contoh, modus dari [10, 4, 5, 2, 4] adalah 4.

Jika modus tidak ditemukan, program akan menghasilkan output -1, apabila ditemukan lebih dari dua nilai modus, tampilkan nilai modus yang paling pertama muncul (dihitung dari kiri ke kanan).

Dan apabila modus hanya memiliki 1 nilai yang sama maka program akan menghasilkan output -1, contohnya [1, 1, 1] adalah -1.

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

### Challenge 6 - Find Majority Element (BONUS)

Buatlah sebuah program yang menamukan elemen mayoritas dari inputan sebuah Array. Elemen mayoritas adalah elemen yang muncul lebih dari n / 2 kali (anda hanya perlu mengeluarkan satu elemen), namun jika tidak ditemukan maka ceta "Tidak Ditemukan".

#### Sample Test Cases
```
Input : []int{3, 3, 4, 2, 4, 4, 2, 4, 4}
Output : 4 

Input : []int{3, 3, 4, 2, 4, 4, 2, 4}
Output : Tidak Ditemukan
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