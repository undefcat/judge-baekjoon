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
	N, M    int
	adj     [][]bool
	nMatch  []int
	mMatch  []int
	visited []int
)

func main() {
	// 직원 N, 일의 갯수 M
	N, M = scanInt(), scanInt()

	adj = make([][]bool, N+1)
	for i := range adj {
		adj[i] = make([]bool, M+1)
	}

	for i := 1; i <= N; i++ {
		jobs := scanInt()
		for ji := 0; ji < jobs; ji++ {
			adj[i][scanInt()] = true
		}
	}

	nMatch = make([]int, N+1)
	mMatch = make([]int, M+1)

	ans := 0
	for n := 1; n <= N; n++ {
		visited = make([]int, N+1)
		if dfs(n) {
			ans++
		}
	}

	fmt.Println(ans)
}

func dfs(n int) bool {
	if visited[n] == 1 {
		return false
	}

	visited[n] = 1

	for m := 1; m <= M; m++ {
		if !adj[n][m] {
			continue
		}

		if mMatch[m] == 0 || dfs(mMatch[m]) {
			nMatch[n] = m
			mMatch[m] = n
			return true
		}
	}

	return false
}