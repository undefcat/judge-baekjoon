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
	plus := make([]bool, 10000001)
	minus := make([]bool, 10000001)

	for i := 0; i < N; i++ {
		num := scanInt()
		if num > 0 {
			plus[num] = true
		} else {
			minus[-num] = true
		}
	}

	M := scanInt()
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	for i := 0; i < M; i++ {
		num := scanInt()
		ans := false
		if num > 0 {
			ans = plus[num]
		} else {
			ans = minus[-num]
		}

		if ans {
			wr.WriteByte('1')
		} else {
			wr.WriteByte('0')
		}
		wr.WriteByte(' ')
	}
}
