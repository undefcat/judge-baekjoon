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
	N, M      int
	adj       [][]bool
	visited   []int
	cowMatch  []int
	barnMatch []int
)

func main() {
	// 소 N, 축사 M
	N, M = scanInt(), scanInt()

	adj = make([][]bool, N+1)
	for i := range adj {
		adj[i] = make([]bool, M+1)
	}

	for cow := 1; cow <= N; cow++ {
		count := scanInt()
		for c := 0; c < count; c++ {
			adj[cow][scanInt()] = true
		}
	}

	cowMatch = make([]int, N+1)
	barnMatch = make([]int, M+1)

	ans := 0
	for cow := 1; cow <= N; cow++ {
		visited = make([]int, N+1)
		if dfs(cow) {
			ans++
		}
	}

	fmt.Println(ans)
}

func dfs(cow int) bool {
	if visited[cow] == 1 {
		return false
	}

	visited[cow] = 1

	for barn := 1; barn <= M; barn++ {
		if !adj[cow][barn] {
			continue
		}

		if barnMatch[barn] == 0 || dfs(barnMatch[barn]) {
			cowMatch[cow] = barn
			barnMatch[barn] = cow
			return true
		}
	}

	return false
}
