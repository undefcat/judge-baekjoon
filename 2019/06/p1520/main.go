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

func scan() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

var (
	M, N     int
	mountain [][]int
	cache [][]int
)

func main() {
	M, N = scan(), scan()
	mountain = make([][]int, M)
	cache = make([][]int, M)

	for i := range mountain {
		mountain[i] = make([]int, N)
		cache[i] = make([]int, N)

		for j := range mountain[i] {
			mountain[i][j] = scan()
			cache[i][j] = -1
		}
	}

	fmt.Print(solve(0, 0))
}

func solve(y, x int) int {
	if y < 0 || x < 0 || y >= M || x >= N {
		return 0
	}

	if (y == M-1) && (x == N-1) {
		return 1
	}

	if cache[y][x] != -1 {
		return cache[y][x]
	}

	ret := 0
	if y+1 < M && (mountain[y][x] > mountain[y+1][x]) {
		ret += solve(y+1, x)
	}

	if y-1 >= 0 && (mountain[y][x] > mountain[y-1][x]) {
		ret += solve(y-1, x)
	}

	if x+1 < N && (mountain[y][x] > mountain[y][x+1]) {
		ret += solve(y, x+1)
	}

	if x-1 >= 0 && (mountain[y][x] > mountain[y][x-1]) {
		ret += solve(y, x-1)
	}

	cache[y][x] = ret
	return ret
}