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

func gridChallenge(grid []string) string {
	for i := 1; i < len(grid); i++ {
		prevLine := []rune(grid[i-1])
		currentLine := []rune(grid[i])
		sort.Slice(prevLine, func(i, j int) bool { return prevLine[i] < prevLine[j] })
		sort.Slice(currentLine, func(i, j int) bool { return currentLine[i] < currentLine[j] })

		for j := 0; j < len(currentLine); j++ {
			if prevLine[j] > currentLine[j] {
				return "NO"
			}
		}
	}
	return "YES"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	stdout, err := os.Create("output.txt")
	checkError(err)
	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)
	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)
		var grid []string
		for i := 0; i < int(n); i++ {
			gridItem := readLine(reader)
			grid = append(grid, gridItem)
		}

		result := gridChallenge(grid)
		fmt.Fprintf(writer, "%s\n", result)
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
