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

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

type byValue []int32

func (f byValue) Len() int {
	return len(f)
}

func (f byValue) Less(i int, j int) bool {
	return f[i] < f[j]
}

func (f byValue) Swap(i int, j int) {
	f[i], f[j] = f[j], f[i]
}

func miniMaxSum(arr []int32) {
	// First sort array
	sort.Sort(byValue(arr))
	var mid, min, max int64
	// Middle three numbers are always included in the sum of 4 elements
	mid = int64(arr[1]) + int64(arr[2]) + int64(arr[3])
	min = mid + int64(arr[0])
	max = mid + int64(arr[4])
	fmt.Printf("%v %v\n", min, max)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
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
