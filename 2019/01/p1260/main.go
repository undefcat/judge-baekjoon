package main

import (
	"bufio"
	"container/list"
	"os"
	"sort"
	"strconv"
)

var (
	sc         = bufio.NewScanner(os.Stdin)
	wr         = bufio.NewWriter(os.Stdout)
	N, M, V    int
	adj        [][]int
)

func makeGraph() {
	N, M, V = scanInt(), scanInt(), scanInt()
	adj = make([][]int, N+1)
	for i := range adj {
		adj[i] = make([]int, 0, 100)
	}

	for i := 0; i < M; i++ {
		u, v := scanInt(), scanInt()
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	for i := 1; i <= N; i++ {
		sort.Ints(adj[i])
	}
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func searchDFS() {
	discovered := make([]bool, N+1)
	DFS(V, discovered)
	wr.WriteByte('\n')
}

func DFS(here int, discovered []bool) {
	discovered[here] = true
	wr.WriteString(strconv.Itoa(here))
	wr.WriteByte(' ')
	for i := range adj[here] {
		there := adj[here][i]
		if !discovered[there] {
			DFS(there, discovered)
		}
	}
}

func searchBFS() {
	queue := list.New()
	discovered := make([]bool, N+1)

	queue.PushBack(V)
	discovered[V] = true
	for queue.Len() != 0 {
		e := queue.Front()
		queue.Remove(e)
		wr.WriteString(strconv.Itoa(e.Value.(int)))
		wr.WriteByte(' ')
		for _, v := range adj[e.Value.(int)] {
			if !discovered[v] {
				queue.PushBack(v)
				discovered[v] = true
			}
		}
	}
	wr.WriteByte('\n')
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	defer wr.Flush()
	makeGraph()
	searchDFS()
	searchBFS()
}
