package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func birthday(s []int, d int, m int) int {
	res := 0
	for i := 0; i <= len(s)-m; i++ {
		sum := 0
		for j := i; j < i+m; j++ {
			sum += s[j]
		}

		if sum == d {
			res += 1
		}
	}

	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int(nTemp)

	sTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var s []int

	for i := 0; i < int(n); i++ {
		sItemTemp, err := strconv.ParseInt(sTemp[i], 10, 64)
		checkError(err)
		sItem := int(sItemTemp)
		s = append(s, sItem)
	}

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	dTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	d := int(dTemp)

	mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	m := int(mTemp)

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
