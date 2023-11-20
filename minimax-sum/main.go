package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func miniMaxSum(arr []int32) {
	// 配列の値は整列されていないため、一番大きな値を取り除いたスライスと、小さな値を取り除いたスライスをsum()にわたす必要がある
	// 方法は以下の2つがありそう
	// 1. 渡された配列をソートしてから、頭の値を取り除いたスライスを計算してmaxの値を出力する。
	//    末尾の値を取り除いたスライスを合計してmini の値を出力する
	//    メリット: 1回のソートで済む
	// 2. 合計を計算する際に、最大の値を検索して取り除き、miniの値を出力する
	//    同様に最小の値を検索して取り除き、maxの値を出力する
	//    max, miniの結果を出力するために2回の検索が必要になる
	intSlice := make([]int, len(arr))
	for i, n := range arr {
		intSlice[i] = int(n)
	}
	sort.Ints(intSlice)
	// quickSort(intSlice)
	fmt.Printf("%v %v\n", sum(intSlice[:len(intSlice)-1]), sum(intSlice[1:]))
}

func sum(arr []int) int {
	var total int
	for _, num := range arr {
		total += num
	}

	return total
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
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
