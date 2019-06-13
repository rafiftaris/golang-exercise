## Alta Go 6 # Data Definition Language

### 🏠 [back to home](https://github.com/alterra-academy/golang-class)

### 🎯 Objektif

- Mampu untuk menggunakan statement SQL DDL
- Mampu menggunakan MySQL di Terminal untuk query data

### ✏️ Challenge

Silakan mempelajari terlebih dahulu lebih lanjut mengenai SQL dan MySQL. Sekarang kamu akan mengimplementasikan schema yang kamu buat menjadi tabel di MySQL, memasukkan data ke tabel dan melakukan pengambilan data pada tabel-tabel yang ada.

### Challenge 1

1. Create database ata_online_shop.
2. Dari schema Olshop yang telah kamu kerjakan di, Implementasikanlah menjadi table pada MySQL.
  a. Create table user.
  b. Create table product, product type, operators, product description, payment_method.
  c. Create table transaction, transaction detail.
3. Create tabel kurir dengan filed id, name, created_at, updated_at.
4. Tambahkan ongkos_dasar column di tabel kurir.
5. Rename tabel kurir menjadi shipping.
6. Hapus / Drop tabel shipping karena ternyata tidak dibutuhkan.
7. Silahkan menambahkan entity baru dengan relation 1-to-1, 1-to-many, many-to-many. Seperti:
  a. 1-to-1: payment method description
  b. 1-to-many: user dengan alamat.
  c. many-to-many: user dengan payment method menjadi user_payment_method_detail.
