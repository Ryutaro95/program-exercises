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
 * Complete the 'matchingStrings' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. STRING_ARRAY strings
 *  2. STRING_ARRAY queries
 */

// func matchingStrings(strings []string, queries []string) []int32 {
// 	// Write your code here
// 	result := make([]int32, len(queries))
// 	for i, val := range queries {
// 		for _, v := range strings {
// 			if val == v {
// 				result[i]++
// 			}
// 		}
// 	}
//
// 	return result
// }

// queriesの文字列をキーとし値をkeyの出現回数としたMapを作成する
// queriesMapのキーと一致する文字列がstrings配列内にあれば、mapの値をインクリメントしていく
// 最終的なアウトプットを作成するために、[]int32を用意して、文字列の搭乗橋回数をまとめたlen(queries)個のスライスを作成してかえす
func matchingStringsHashMapVer(strings []string, queries []string) []int32 {
	queriesMap := createmapFromArray(queries)

	for _, str := range strings {
		if _, ok := queriesMap[str]; ok {
			queriesMap[str]++
		}
	}

	var result []int32
	for _, query := range queries {
		result = append(result, queriesMap[query])
	}
	return result
}

func createmapFromArray(queries []string) map[string]int32 {
	result := make(map[string]int32)

	for _, val := range queries {
		result[val] = 0
	}
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create("output.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	stringsCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var strs []string

	for i := 0; i < int(stringsCount); i++ {
		stringsItem := readLine(reader)
		strs = append(strs, stringsItem)
	}

	queriesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var queries []string

	for i := 0; i < int(queriesCount); i++ {
		queriesItem := readLine(reader)
		queries = append(queries, queriesItem)
	}

	res := matchingStringsHashMapVer(strs, queries)

	for i, resItem := range res {
		fmt.Fprintf(writer, "%d", resItem)

		if i != len(res)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
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
