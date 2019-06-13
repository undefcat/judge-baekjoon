package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	sc *bufio.Scanner
	wr *bufio.Writer
)
func init() {
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
	sc.Split(bufio.ScanWords)
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

var (
	N     int
	A     []int
	cache []int
)

func main() {
	defer wr.Flush()

	N = scanInt()
	A = make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = scanInt()
	}

	cache = make([]int, N)

	maxVal := 0
	maxIdx := 0
	for i := 0; i < N; i++ {
		val := lis(i)
		if val > maxVal {
			maxVal = val
			maxIdx = i
		}
	}

	wr.WriteString(strconv.Itoa(maxVal))
	wr.WriteByte('\n')
	reconstruct(maxIdx)
}

func lis(n int) int {
	if n == N {
		return 1
	}

	ret := &cache[n]
	if *ret != 0 {
		return *ret
	}

	*ret = 1
	for next := n+1; next < N; next++ {
		if A[n] < A[next] {
			*ret = max(*ret, lis(next)+1)
		}
	}

	return *ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func reconstruct(n int) {
	if lis(n) == 1 {
		wr.WriteString(strconv.Itoa(A[n]))
		return
	}

	for next := n+1; next < N; next++ {
		if lis(n) == lis(next)+1 {
			wr.WriteString(strconv.Itoa(A[n]))
			wr.WriteByte(' ')
			reconstruct(next)
			return
		}
	}
}
