package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func alternate(s string) int {
	set := make(map[rune]int)
	for _, c := range s {
		set[c] += 1
	}

	keys := make([]rune, len(set))
	i := 0
	for k := range set {
		keys[i] = k
		i++
	}

	max := 0
	for i := 0; i < len(keys)-1; i++ {
		for j := i + 1; j < len(keys); j++ {
			res := isAlternate(s, keys[i], keys[j])
			if res {
				newLen := set[keys[i]] + set[keys[j]]
				if newLen > max {
					max = newLen
				}
			}
		}
	}

	return max
}

func isAlternate(s string, a, b rune) bool {
	before := ' '

	for _, c := range s {
		if c == a || c == b {
			if before == c {
				return false
			}
			before = c
		}
	}

	return true
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	_, err = strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	//l := int32(lTemp)

	s := readLine(reader)

	result := alternate(s)

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
