package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func timeConversion(s string) string {
	modifier := s[8:]
	hh := s[:2]

	hhInt, err := strconv.Atoi(hh)
	if err != nil {
		panic(err)
	}

	var correctTime string
	if modifier == "PM" && hhInt != 12 {
		correctTime = strconv.Itoa(hhInt + 12)
	} else if modifier == "AM" && hhInt == 12 {
		correctTime = "00"
	} else {
		return s[:8]
	}
	return fmt.Sprintf("%s\n", strings.Join([]string{correctTime, s[2:8]}, ""))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create("./output.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)
	fmt.Println(result)

	fmt.Fprintf(writer, "%s\n", result)

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
		fmt.Println(err)
		log.Fatal(err)
		panic(err)
	}
}
