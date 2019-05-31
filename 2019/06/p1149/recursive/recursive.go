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
	cost  [][3]int
	cache [][3]int
	tc    int
)

func main() {
	tc = scan()
	cost = make([][3]int, tc)
	cache = make([][3]int, tc)

	for i := 0; i < tc; i++ {
		cost[i][0], cost[i][1], cost[i][2] = scan(), scan(), scan()
	}

	ans := math.MaxInt64
	for color := 0; color < 3; color++ {
		ans = min(ans, solve(0, color))
	}

	fmt.Print(ans)
}

func solve(i, color int) int {
	if i == tc {
		return 0
	}

	if cache[i][color] != 0 {
		return cache[i][color]
	}

	minValue := math.MaxInt32

	for j := 0; j < 3; j++ {
		if color == j {
			continue
		}
		minValue = min(minValue, solve(i+1, j))
	}

	cache[i][color] = minValue + cost[i][color]
	return cache[i][color]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
