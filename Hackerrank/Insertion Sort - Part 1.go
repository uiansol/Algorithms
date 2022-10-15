package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func insertionSort1(n int32, arr []int32) {
	i := n - 1
	toAdd := arr[i]
	for {
		if i == 0 || arr[i-1] < toAdd {
			break
		}
		arr[i] = arr[i-1]
		temp := fmt.Sprintf("%v", arr)
		fmt.Println(strings.Join(strings.Split(temp[1:len(temp)-1], " "), " "))
		i -= 1
	}

	arr[i] = toAdd
	temp := fmt.Sprintf("%v", arr)
	fmt.Println(strings.Join(strings.Split(temp[1:len(temp)-1], " "), " "))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	insertionSort1(n, arr)
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
