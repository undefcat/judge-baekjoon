package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	sc         = bufio.NewScanner(os.Stdin)
	N, M       int
	adj        [][]int
	discovered []int
	sccGroup   []int
	stack      []int
	top        = -1
	counter    = 0
	sccId      = 1
)

func scanInt() int {
	sc.Scan()
	input := sc.Bytes()
	neg := 1
	if input[0] == '-' {
		neg = -1
		input = input[1:]
	}

	sum := 0
	for _, v := range input {
		sum *= 10
		sum += int(v - '0')
	}

	return neg * sum
}

func makeGraph() {
	N, M = scanInt(), scanInt()
	adj = make([][]int, 2*N)

	for i := range adj {
		adj[i] = make([]int, 0, 10)
	}

	for i := 0; i < M; i++ {
		a, b := scanInt(), scanInt()
		na, nb := con(-a), con(-b)
		// !a -> b
		// !b -> a
		adj[na] = append(adj[na], con(b))
		adj[nb] = append(adj[nb], con(a))
	}
}

func tarjanSCC() {
	discovered = make([]int, 2*N)
	sccGroup = make([]int, 2*N)
	stack = make([]int, 2*N)

	for i := 0; i < 2*N; i++ {
		if discovered[i] == 0 {
			scc(i)
		}
	}
}

func scc(here int) int {
	counter++
	discovered[here] = counter
	ret := discovered[here]
	top++
	stack[top] = here

	for i := 0; i < len(adj[here]); i++ {
		there := adj[here][i]
		if discovered[there] == 0 {
			ret = min(ret, scc(there))
		} else if sccGroup[there] == 0 {
			ret = min(ret, discovered[there])
		}
	}

	if ret == discovered[here] {
		for {
			variable := stack[top]
			top--
			sccGroup[variable] = sccId
			if variable == here {
				break
			}
		}
		sccId++
	}

	return ret
}

func con(n int) int {
	if n < 0 {
		return -2*(n+1)+1
	}
	return 2*(n-1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func canSolved() bool {
	for i := 0; i < 2*N; i += 2 {
		if sccGroup[i] == sccGroup[i+1] {
			return false
		}
	}
	return true
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	makeGraph()
	tarjanSCC()
	if !canSolved() {
		fmt.Print(0)
		return
	}
	fmt.Print(1)
}
