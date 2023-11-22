package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func flippingBits(n int64) int64 {
	binary := decimalToBinaryString(n)
	converted := ""
	for _, bin := range binary {
		if bin == '0' {
			converted = fmt.Sprintf("%s%d", converted, 1)
		} else {
			converted = fmt.Sprintf("%s%d", converted, 0)
		}
	}

	result, err := strconv.ParseUint(converted, 2, 32)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T, %d\n", result, result)
	return int64(result)
}

func decimalToBinaryString(decimal int64) string {
	if decimal == 0 {
		return "00000000000000000000000000000000"
	}
	binary := ""

	for decimal > 0 {
		remainder := decimal % 2
		binary = fmt.Sprintf("%d%s", remainder, binary)
		decimal /= 2
	}

	padding := 32 - len(binary)
	for i := 0; i < padding; i++ {
		binary = fmt.Sprintf("%d%s", 0, binary)
	}
	return binary
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create("output.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)
	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)

		result := flippingBits(n)
		fmt.Fprintf(writer, "%d\n", result)
	}

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
