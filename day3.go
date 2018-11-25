package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	N := int32(NTemp)
	evenOdd(N)

}

// If n is odd, print "Weird"
// If n is even and in the inclusive range of 2 to 5, print "Not Weird"
// If n is even and in the inclusive range of to 6, 20 print "Weird"
// If n is even and greater than 20, print "Not Weird"

func evenOdd(num int32) {
	if num%2 != 0 {
		fmt.Println("Weird")
	} else if num%2 == 0 && (num > 1 && num < 6) {
		fmt.Println("Not Weird")
	} else if num%2 == 0 && (num > 5 && num < 21) {
		fmt.Println("Weird")
	} else {
		fmt.Println("Not Weird")
	}
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
