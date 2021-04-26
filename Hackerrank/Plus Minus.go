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
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int) {
	pos, neg, zer := 0, 0, 0
	n := len(arr)

	for i := 0; i < n; i++ {
		if arr[i] > 0 {
			pos += 1
		} else if arr[i] < 0 {
			neg += 1
		} else {
			zer += 1
		}
	}

	s := fmt.Sprintf("%.6f\n", float64(pos)/float64(n))
	fmt.Printf(s)
	s = fmt.Sprintf("%.6f\n", float64(neg)/float64(n))
	fmt.Printf(s)
	s = fmt.Sprintf("%.6f\n", float64(zer)/float64(n))
	fmt.Printf(s)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
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
