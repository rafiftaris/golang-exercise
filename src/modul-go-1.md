# Mengakses Nilai Dalam Array

Tahukah kamu jika *string* adalah sebuah *array*? Kamu dapat mengakses karakter-karakter pada sebuah string layaknya mengakses nilai pada sebuah array. Untuk membuktikannya, kerjakanlah tantangan ini!

## Objectives

- Mengetahui *properties* Pada Array
- Mengerti Cara Mengakses Nilai Dalam Array

## Directions

Buatlah sebuah fungsi dengan nama **balikString**. Fungsi ini akan menerima argumen sebuah string dan mengembalikan kebalikannya. 


### Tugas 1

Ubahlah seluruh `var` pada kode JavaScript dibawah ini menjadi `let` atau `const`!

```javascript
var phi = 3.14;
var power = 2;
var radius = 0;
var calculateArea = function (radius) {
  return phi * Math.pow(radius, power);
}

radius = 21;
var area21 = calculateArea(radius);

radius = 7;
var area7 = calculateArea(radius);

console.log('area with 21 radius: ' + area21 + ', and area with 7 radius: ' + area7);
```

### Tugas 2

Ubahlah anonymous functions dibawah ini menjadi ES6 Arrow Function!

```javascript
  const multiply = function (num1, num2) {
    return num1 * num2;
  }

  const divide = function (num1, num2) {
    return num1 / num2;
  }

  const doubleMe = function (num) {
    return num * 2;
  }

  console.log(multiply(5,2));
  console.log(divide(10,2));
```