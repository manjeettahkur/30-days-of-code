package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	counter := 0
	var n int
	person := make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()
		if counter == 0 {
			counter ++
			n, _ = strconv.Atoi(line)
		} else if counter <= n {
			counter ++
			temp := strings.Split(line, " ")
			person[temp[0]] = temp[1]

		} else {
			if val, ok := person[line]; ok {
				fmt.Printf("%s=%s\n",line,val)
			} else {
				fmt.Println("Not found")
			}
		}
	}
}
