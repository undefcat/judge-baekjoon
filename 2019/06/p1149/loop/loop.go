package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
}

func scan() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

var (
	cost [][3]int
	cache [][3]int
)

func main() {
	tc := scan()
	cost = make([][3]int, tc)
	cache = make([][3]int, tc)

	for i := 0; i < tc; i++ {
		cost[i][0], cost[i][1], cost[i][2] = scan(), scan(), scan()
	}

	for i := 0; i < 3; i++ {
		cache[0][i] = cost[0][i]
	}

	for i := 1; i < tc; i++ {
		for j := 0; j < 3; j++ {
			cache[i][j] = min(cache[i-1][(j+1)%3], cache[i-1][(j+2)%3])+cost[i][j]
		}
	}

	ans := math.MaxInt64
	for i := 0; i < 3; i++ {
		ans = min(ans, cache[tc-1][i])
	}
	fmt.Print(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
