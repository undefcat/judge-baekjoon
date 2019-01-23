package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc         = bufio.NewScanner(os.Stdin)
	wr         = bufio.NewWriter(os.Stdout)
	V, C       int
	adj        [][]int
	discovered []int
	sccGroup   []int
	counter    int
	sccId      = 1
)

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func makeGraph() {
	V, C = scanInt(), scanInt()
	adj = make([][]int, 2*V)

	for i := range adj {
		adj[i] = make([]int, 0, 10)
	}

	// 입력 받을 때에만 1, 2, 3.. 변수로 생각한다.
	// 그 외에 내가 문제를 풀 때에는 0, 1, 2 로 생각한다.
	// !a => b
	// !b => a
	for i := 0; i < C; i++ {
		a, b := scanInt(), scanInt()
		adj[con(-a)] = append(adj[con(-a)], con(b))
		adj[con(-b)] = append(adj[con(-b)], con(a))
	}
}

func con(n int) int {
	if n > 0 {
		return 2*(n-1)
	}

	return -2*(n+1)+1
}

var (
	stack []int
	top   = -1
)

func makeSCC() {
	stack = make([]int, 2*V)
	sccGroup = make([]int, 2*V)
	discovered = make([]int, 2*V)

	for i := range adj {
		if discovered[i] == 0 {
			scc(i)
		}
	}
}

func scc(here int) (ret int) {
	counter++
	discovered[here] = counter
	top++
	stack[top] = here
	ret = discovered[here]

	for i := range adj[here] {
		there := adj[here][i]
		if discovered[there] == 0 {
			ret = min(ret, scc(there))
		} else if sccGroup[there] == 0 {
			ret = min(ret, discovered[there])
		}
	}

	if ret == discovered[here] {
		for {
			value := stack[top]
			top--
			sccGroup[value] = sccId
			if value == here {
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

func canSolved() bool {
	for i := 0; i < 2*V; i += 2 {
		if sccGroup[i] == sccGroup[i+1] {
			return false
		}
	}
	return true
}

func solved() {
	for i := 0; i < V*2; i += 2 {
		// !xi -> xi이면 xi는 참
		if sccGroup[i] < sccGroup[i+1] {
			wr.WriteByte('1')
		} else {
			wr.WriteByte('0')
		}
		wr.WriteByte(' ')
	}

	wr.Flush()
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	makeGraph()
	makeSCC()
	if !canSolved() {
		fmt.Print(0)
		return
	}

	fmt.Println(1)
	solved()
}
