package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the staircase function below.
func staircase(n int32) {
	fmt.Printf(makeStair(n))
}

func makeStair(n int32) string {
	var stair string
	for i := 1; i <= int(n); i++ {
		stair += strings.Repeat(" ", int(n)-i) + strings.Repeat("#", i)
		if i < int(n) {
			stair += "\n"
		}
	}
	return stair
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	staircase(n)
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
