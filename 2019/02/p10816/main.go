package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	sc *bufio.Scanner
)

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	N := scanInt()
	plus := make([]int, 10000001)
	minus := make([]int, 10000001)

	for i := 0; i < N; i++ {
		num := scanInt()
		if num > 0 {
			plus[num]++
		} else {
			minus[-num]++
		}
	}

	M := scanInt()
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	for i := 0; i < M; i++ {
		num := scanInt()

		if num > 0 {
			wr.WriteString(strconv.Itoa(plus[num]))
		} else {
			wr.WriteString(strconv.Itoa(minus[-num]))
		}

		wr.WriteByte(' ')
	}
}
