# Palindrome Number

整数 `x` が与えられたとき、 `x` が回文であれば `true` を返す

## Example
1: 
```go
Input: x = 121
Output: true
```

2: 
```go
Input: x = -121
Output: false
```
3:
```go
Input: x = 10
Output: false
```

## Approach
整数配列 x を文字列に変換した後に両端を順に比較していき、両端のindex番号が等しくなったら処理を抜ける。等しくなければ `false` を返す
```go
step1:

        i                 (len-1)-i
str := [1, 2, 3, 4, 3, 2, 1]
1 == 1

step2
           i        (len-1)-i
           ↓           ↓
str := [1, 2, 3, 4, 3, 2, 1]
2 == 2

step3
              i   (len-1)-i
              ↓     ↓
str := [1, 2, 3, 4, 3, 2, 1]
3 == 3

step4
                 i
             (len-1)-i
                 ↓ 
str := [1, 2, 3, 4, 3, 2, 1]
両端のindex番号が等しくなったので、処理を抜ける

```

## Leetcode Editorial
上の実装では、整数を文字列に変換しているが、これは文字列を作成する他の余分な非定数スペースが許可されていないので別のアプローチを取るべきらしい

別のアプローチとして、数値を半分にして後半を反転した後に前半と後半を比較する。  
inputが `1221` の場合、 `12` と `21` に分けられる。 そして `21` を `12` に反転する。これで 前半と後半を比較することができるようになった

また、負の数は `-` が付くため、全て回文とはならない。なのですべて `false` 返すことができる。

o
```go
x = 1234567654321

x: 12345676543, revertedNumber: 13
x: 1234567654,  revertedNumber: 123
x: 123456765,   revertedNumber: 1234
x: 12345676,    revertedNumber: 12345
x: 1234567,     revertedNumber: 123456
x: 123456,      revertedNumber: 1234567

revertedNumber := 0
for x > revertedNumber {
    revertedNumber = revertedNumber*10 + x%10
    x /= 10
}

```