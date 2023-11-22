package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func stringsXOR(s, t string) string {
	var answer string
	for i := range s {
		if s[i] == t[i] {
			answer = fmt.Sprintf("%s%s", answer, "0")
		} else {
			answer = fmt.Sprintf("%s%s", answer, "1")
		}
	}
	return answer
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	sTmp := strings.TrimSpace(readLine(reader))
	tTmp := strings.TrimSpace(readLine(reader))

	fmt.Println(stringsXOR(sTmp, tTmp))
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}
