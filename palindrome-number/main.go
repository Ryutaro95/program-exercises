package main

import "fmt"

func isPalindrome(x int) bool {
	// 特別なケース
	// * -121 などの負の値は回文にならないので false
	// * 末尾が0の場合、回文となるためには頭も0である必要がある
	// * ただ、この条件が当てはまるのは 0 のみなのでそれ以外は false
	// x%10 で一桁目を取得
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10

		x /= 10

		fmt.Printf("x: %d, revertedNumber: %d\n", x, revertedNumber)
		/*
		   x: 12345676543, revertedNumber: 13
		   x: 1234567654,  revertedNumber: 123
		   x: 123456765,   revertedNumber: 1234
		   x: 12345676,    revertedNumber: 12345
		   x: 1234567,     revertedNumber: 123456
		   x: 123456,      revertedNumber: 1234567
		*/
	}

	if x != revertedNumber && x != revertedNumber/10 {
		return false
	}

	return true
}

func main() {
}

/*
   *   int revertedNumber = 0;
       while(x > revertedNumber) {
           revertedNumber = revertedNumber * 10 + x % 10;
           x /= 10;
       }
*/
