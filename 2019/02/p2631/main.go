package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	sc *bufio.Scanner
)

func scanInt() int {
	sc.Scan()
	num := 0
	for _, v := range sc.Bytes() {
		num *= 10
		num += int(v-'0')
	}
	return num
}

func init() {
	sc = bufio.NewScanner(os.Stdin)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	N := scanInt()
	arr := make([]int, N)
	dp := make([]int, N)
	dp[0] = 1
	for i := 1; i < N; i++ {
		copy(dp[i:], dp[:i])
	}

	for i := 0; i < N; i++ {
		arr[i] = scanInt()
	}

	for i := 1; i < N; i++ {
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	lcs := 0
	for _, v := range dp {
		lcs = max(lcs, v)
	}

	fmt.Print(N-lcs)
}
