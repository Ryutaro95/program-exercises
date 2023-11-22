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
* Exmaple Input:
  31
  2 2 2 1 3 2 2 3 3 1 4 1 3 2 2 1 2 2 4 2 2 3 5 3 4 3 2 1 4 5 4
  10 4
* Example Output:
  7
*/

func birthday(s []int32, d, m int32) int32 {
	var answer int32
	for i := range s {
		if (int(m) + i) > len(s) {
			break
		}
		dividedArr := s[i : int(m)+i]

		var sum int32
		for _, item := range dividedArr {
			sum += item
		}
		if sum == d {
			answer++
		}
	}

	return answer
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	stdout, err := os.Create("output.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)
	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	sTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	var s []int32

	for i := 0; i < int(n); i++ {
		sItemTemp, err := strconv.ParseInt(sTemp[i], 10, 64)
		checkError(err)
		sItem := int32(sItemTemp)
		s = append(s, sItem)
	}

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	dTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	d := int32(dTemp)

	mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	result := birthday(s, d, m)

	fmt.Fprintf(writer, "%d\n", result)
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
