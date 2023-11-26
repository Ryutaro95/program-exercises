package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func BalancedSums(arr []int32) string {
	var total int32
	for _, val := range arr {
		total += val
	}

	var leftSum int32
	for i := 0; i < len(arr); i++ {
		rightSum := total - leftSum - arr[i]
		if leftSum == rightSum {
			return "YES"
		}
		leftSum += arr[i]
	}
	return "NO"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	testCaseCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	stdout, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer stdout.Close()
	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	for t := 0; t < int(testCaseCount); t++ {
		nInt, err := strconv.ParseInt(readLine(reader), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		n := int32(nInt)
		temp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
		var arr []int32

		for i := 0; i < int(n); i++ {
			item, err := strconv.ParseInt(temp[i], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			arrItem := int32(item)
			arr = append(arr, arrItem)
		}

		result := BalancedSums(arr)
		fmt.Fprintf(writer, "%s\n", result)
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(string(str))
}
