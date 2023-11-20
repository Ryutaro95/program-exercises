package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
	// Write your code here
	var plusCount, minusCount, zeroCount int

	for _, v := range arr {
		if v < 0 {
			minusCount++
		} else if v > 0 {
			plusCount++
		} else {
			zeroCount++
		}
	}
	n := float64(len(arr))
	fmt.Printf("%.6f\n%.6f\n%.6f", float64(plusCount)/n, float64(minusCount)/n, float64(zeroCount)/n)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	// strconv.ParseInt(s string, base int, bitSize int) (int64, error)
	// 第一引数で与えた文字列を整数(int64)に変化して返す関数
	// 2引数: 整数の基数(10進数 or 16進数 or 8進数)
	// 3引数: 結果の整数型のビットサイズ(32 or 64)
	//
	// 1行名のstdinで渡された値を取得
	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	// 2行目のstdinで渡された値を取得。今回の場合は、配列なる整数の塊
	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	// arrTemp: []string

	var arr []int32

	// []string -> []int32へ変換している
	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
