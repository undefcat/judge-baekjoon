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
	sc.Split(bufio.ScanWords)

	wr = bufio.NewWriter(os.Stdout)
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func main() {
	defer wr.Flush()

	tc := scanInt()
	for tci := 0; tci < tc; tci++ {
		l, n := scanInt(), scanInt()
		mid := l/2

		fast, slow := l+1, 0
		closestMid := 1000001
		for ni := 0; ni < n; ni++ {
			ant := scanInt()
			if ant == 0 || ant == l {
				// 솔직히 개미들이 막대 끝에만 있으면
				// 가장 느린값은 0이어야 하지 않나?
				slow = l
			}

			if ant > mid {
				if closestMid > ant-mid {
					closestMid = ant-mid
					fast = l-ant
				}
			} else {
				if closestMid > mid-ant {
					closestMid = mid-ant
					fast = ant
				}
			}

			slow = max(slow, l-ant)
			slow = max(slow, ant)
		}

		if fast == l+1 {
			fast = 0
		}

		wr.WriteString(strconv.Itoa(fast))
		wr.WriteByte(' ')
		wr.WriteString(strconv.Itoa(slow))
		wr.WriteByte('\n')
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}