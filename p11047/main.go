package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
)

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	N, K := scanInt(), scanInt()
	coins := make([]int, N)

	for i := N-1; i >= 0; i-- {
		coins[i] = scanInt()
	}

	ans := 0
	for i := range coins {
		for {
			if K >= coins[i] {
				K -= coins[i]
				ans++
			} else {
				break
			}
		}
	}

	fmt.Print(ans)
}
