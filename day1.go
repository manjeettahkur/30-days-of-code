package main


import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d = 4.0
	var s = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	for j := 0; j <= 3 && scanner.Scan(); j++ {
		switch j {
		case 0:
			num, _ := strconv.Atoi(scanner.Text())
			i += uint64(num)
			fmt.Println(i)
		case 1:
			num, _ := strconv.ParseFloat(scanner.Text(), 64)
			d += num
			s := fmt.Sprintf("%.1f", d)
			fmt.Println(s)
		case 2:
			s += scanner.Text()
			fmt.Println(s)

		}
	}
}