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
 * Complete the 'weightedUniformStrings' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. INTEGER_ARRAY queries
 */

func weightedUniformStrings(s string, queries []int) []string {
	solutions := make(map[int]bool)
	solutions[int(s[0]-'a'+1)] = true
	n, previous, repeat := len(s), 0, 1

	for i := 1; i < n; i++ {
		idx := int(s[i] - 'a' + 1)
		if s[i] != s[previous] {
			repeat = 1
			solutions[idx] = true
		} else {
			repeat++
			solutions[idx*repeat] = true
		}
		previous = i
	}

	result := []string{}
	for _, v := range queries {
		if _, ok := solutions[v]; ok {
			result = append(result, "Yes")
		} else {
			result = append(result, "No")
		}
	}

	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	queriesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var queries []int

	for i := 0; i < int(queriesCount); i++ {
		queriesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		queriesItem := int(queriesItemTemp)
		queries = append(queries, queriesItem)
	}

	result := weightedUniformStrings(s, queries)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

		if i != len(result)-1 {
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
