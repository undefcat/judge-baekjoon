package rank1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc  *bufio.Scanner
	N   int
	D   [][]byte
	ans []int
)

func init() {
	sc = bufio.NewScanner(os.Stdin)
}

func scanString() string {
	sc.Scan()
	return sc.Text()
}

func scanInt() int {
	n, _ := strconv.Atoi(scanString())
	return n
}

func main() {
	N = scanInt()
	D = make([][]byte, N+1)
	for di := range D {
		D[di] = make([]byte, N+1)
	}

	ans = make([]int, N+1)

	str := scanString()
	for i := 1; i <= N; i++ {
		for j := i; j <= N; j++ {
			D[i][j] = str[0]
			str = str[1:]
		}
	}

	dfs(1)
}

func dfs(x int) bool {
	if x == N+1 {
		for _, v := range ans[1:N+1] {
			fmt.Printf("%d ", v)
		}
		return true
	}

	mn := -10
	mx := 10
	t := 0

	for i := x; i >= 1; i-- {
		var start, end int

		switch D[i][x] {
		case '+':
			start, end = 1, 10

		case '-':
			start, end = -10, -1
		}

		start -= t
		end -= t

		mn = max(mn, start)
		mx = min(mx, end)

		t += ans[i-1]
	}

	for ans[x] = mn; ans[x] <= mx; ans[x]++ {
		if dfs(x+1) {
			return true
		}
	}

	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}