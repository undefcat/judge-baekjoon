package main

import (
	"bufio"
	"container/list"
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

func main() {
	// 소 N, 축사 M
	N, M := scanInt(), scanInt()

	graph := make([][]int, N)
	for i := range graph {
		graph[i] = make([]int, 0, M)
	}

	for n := 0; n < N; n++ {
		m := scanInt()
		for mi := 0; mi < m; mi++ {
			graph[n] = append(graph[n], scanInt()-1)
		}
	}

	fmt.Println(networkFlow(graph, N, M))
}

// src 정점은 0
// sink 정점은 1
func networkFlow(graph [][]int, cow, cage int) int {
	V := cow+cage+2

	capacity := make([][]int, V)
	flow := make([][]int, V)

	for i := range capacity {
		capacity[i] = make([]int, V)
		flow[i] = make([]int, V)
	}

	// src와 소를 연결한다.
	for c := 2; c < 2+cow; c++ {
		capacity[0][c] = 1

		// 소와 축사를 연결한다.
		for _, cage := range graph[c-2] {
			capacity[c][cage+2+cow] = 1
		}
	}

	// 축사와 sink를 연결한다.
	for cage := 2+cow; cage < V; cage++ {
		capacity[cage][1] = 1
	}

	totalFlow := 0
	for {
		parent := make([]int, V)
		parent[0] = -1
		for pi := 1; pi < V; pi *= 2 {
			copy(parent[pi:], parent[:pi])
		}
		parent[0] = 0

		q := list.New()
		q.PushBack(0)

		for q.Len() != 0 && parent[1] == -1 {
			e := q.Front()
			q.Remove(e)
			here := e.Value.(int)

			for there := 0; there < V; there++ {
				// 잔여 용량이 있는 곳 탐색
				if capacity[here][there] - flow[here][there] > 0 &&
					parent[there] == -1 {

					q.PushBack(there)
					parent[there] = here
				}
			}
		}

		if parent[1] == -1 {
			break
		}

		// 최대로 보낼 수 있는 유량이 얼마인지 찾는다.
		amount := 987654321
		for p := 1; p != 0; p = parent[p] {
			amount = min(capacity[parent[p]][p] - flow[parent[p]][p],
				amount)
		}

		for p := 1; p != 0; p = parent[p] {
			flow[parent[p]][p] += amount
			flow[p][parent[p]] -= amount
		}

		totalFlow += amount
	}

	return totalFlow
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}