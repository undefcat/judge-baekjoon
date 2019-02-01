package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

var adj [][]int

func makeGraph() {
	N, M := scanInt(), scanInt()
	adj = make([][]int, N*2+1)
	for i := range adj {
		adj[i] = make([]int, 0, N)
	}

	for i := 0; i < M; i++ {
		a, b := scanInt(), scanInt()
		adj[con(-b)] = append(adj[con(-b)], con(a))
		adj[con(-a)] = append(adj[con(-a)], con(b))
	}
}

func con(n int) int {
	if n > 0 {
		return 2*n-1
	}
	return -2*n
}

var (
	count    int
	sccId    int
	visit    []int
	sccGroup []int
	stack    = list.New()
)

func makeTarjanScc() {
	visit = make([]int, len(adj))
	sccGroup = make([]int, len(adj))
	for i := 0; i < len(adj); i++ {
		if visit[i] == 0 {
			scc(i)
		}
	}
}

func scc(here int) int {
	count++
	visit[here] = count
	stack.PushBack(here)
	ret := visit[here]

	for j := 0; j < len(adj[here]); j++ {
		there := adj[here][j]

		if visit[there] == 0 {
			ret = min(ret, scc(there))
		} else if sccGroup[there] == 0 {
			ret = min(ret, visit[there])
		}
	}

	if ret == visit[here] {
		for {
			e := stack.Back()
			variable := e.Value.(int)
			stack.Remove(e)
			sccGroup[variable] = sccId
			if variable == here {
				break
			}
		}
		sccId++
	}
	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	makeGraph()
	makeTarjanScc()

	for i := 1; i < len(adj)-1; i += 2 {
		if sccGroup[i] == sccGroup[i+1] {
			fmt.Print(0)
			return
		}
	}

	fmt.Print(1)
}
