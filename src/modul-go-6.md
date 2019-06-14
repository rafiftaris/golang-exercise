## Alta Go 6 # Data Definition Language

### 🏠 [back to home](https://github.com/alterra-academy/golang-class)

### 🎯 Objektif

<<<<<<< HEAD
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
=======
- Mampu untuk menggunakan statement SQL DDL
- Mampu menggunakan MySQL di Terminal untuk query data

### ✏️ Challenge

Silakan mempelajari terlebih dahulu lebih lanjut mengenai SQL dan MySQL. Sekarang kamu akan mengimplementasikan schema yang kamu buat menjadi tabel di MySQL, memasukkan data ke tabel dan melakukan pengambilan data pada tabel-tabel yang ada.

### Challenge 1

1. Create database ata_online_shop.
2. Dari schema Olshop yang telah kamu kerjakan di, Implementasikanlah menjadi table pada MySQL.
    * Create table user.
    * Create table product, product type, operators, product description, payment_method.
    * Create table transaction, transaction detail.
3. Create tabel kurir dengan filed id, name, created_at, updated_at.
4. Tambahkan ongkos_dasar column di tabel kurir.
5. Rename tabel kurir menjadi shipping.
6. Hapus / Drop tabel shipping karena ternyata tidak dibutuhkan.
7. Silahkan menambahkan entity baru dengan relation 1-to-1, 1-to-many, many-to-many. Seperti:
    * 1-to-1: payment method description
    * 1-to-many: user dengan alamat.
    * many-to-many: user dengan payment method menjadi user_payment_method_detail.
>>>>>>> bd718adde3da3ba2719fc77d909865e9fc7a38bc
