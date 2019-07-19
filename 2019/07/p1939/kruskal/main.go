package kruskal

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type UFSet struct {
	parent []int
	rank   []int
}

func NewUFSet(size int) *UFSet {
	parent := make([]int, size)
	for i := range parent {
		parent[i] = i
	}

	return &UFSet{
		parent: parent,
		rank: make([]int, size)}
}

func (it *UFSet) Find(u int) int {
	p := &it.parent[u]
	if *p == u {
		return u
	}

	*p = it.Find(*p)
	return *p
}

func (it *UFSet) Union(u, v int) int {
	u, v = it.Find(u), it.Find(v)

	if u == v {
		return u
	}

	if it.rank[u] > it.rank[v] {
		u, v = v, u
	}

	if it.rank[u] == it.rank[v] {
		it.rank[v]++
	}

	it.parent[u] = v
	return v
}

type Bridge struct {
	weight int
	a, b int
}

var (
	N, M int
)

func main() {
	N, M = scanInt(), scanInt()

	bridges := make([]*Bridge, M)
	for mi := 0; mi < M; mi++ {
		a, b, weight := scanInt()-1, scanInt()-1, scanInt()
		bridges[mi] = &Bridge{weight, a, b}
	}

	sort.Slice(bridges, func(i, j int) bool {
		return bridges[i].weight > bridges[j].weight
	})

	from, to := scanInt()-1, scanInt()-1
	uf := NewUFSet(N)

	for _, bridge := range bridges {
		uf.Union(bridge.a, bridge.b)

		if uf.Find(from) == uf.Find(to) {
			fmt.Print(bridge.weight)
			return
		}
	}
}