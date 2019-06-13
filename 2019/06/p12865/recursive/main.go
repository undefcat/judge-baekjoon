package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func init() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

var (
	N, K  int
	W     []int
	V     []int
	cache [][]int
)

func main() {
	N, K = scanInt(), scanInt()
	W = make([]int, N)
	V = make([]int, N)

	cache = make([][]int, N)
	for ci := range cache {
		cache[ci] = make([]int, K+1)
		for cj := range cache[ci] {
			cache[ci][cj] = -1
		}
	}

	for i := 0; i < N; i++ {
		W[i] = scanInt()
		V[i] = scanInt()
	}

	fmt.Print(happy(0, K))
}

func happy(n, weight int) int {
	if n == N {
		return 0
	}

	ret := &cache[n][weight]
	if *ret != -1 {
		return *ret
	}

	*ret = 0
	if weight-W[n] >= 0 {
		*ret = happy(n+1, weight-W[n]) + V[n]
	}
	*ret = max(*ret, happy(n+1, weight))

	return *ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
