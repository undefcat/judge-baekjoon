package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

var (
	sc         = bufio.NewScanner(os.Stdin)
	V, E       int
	adj        [][]int
	discovered []int
	counter    int
	sccGroups  groups
	inSCC      []bool
)

type group struct {
	min        int
	vertexList []int
}

func (it *group) add(vertex int) {
	it.vertexList = append(it.vertexList, vertex)
	it.min = min(it.min, vertex)
}

type groups []*group

func (it groups) Len() int {
	return len(it)
}

func (it groups) Less(i, j int) bool {
	return it[i].min < it[j].min
}

func (it groups) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func makeGraph() {
	V, E = scanInt(), scanInt()
	adj = make([][]int, V+1)
	discovered = make([]int, V+1)
	inSCC = make([]bool, V+1)
	sccGroups = make(groups, 0, 100)

	for i := 0; i <= V; i++ {
		adj[i] = make([]int, 0, 10)
	}

	for i := 1; i <= E; i++ {
		u, v := scanInt(), scanInt()
		adj[u] = append(adj[u], v)
	}
}

func doTarjanSCC() {
	for i := 1; i <= V; i++ {
		if discovered[i] == 0 {
			tarjanSCC(i)
		}
	}
}

var (
	stack = make([]int, 10000)
	top   = -1
)

func tarjanSCC(here int) int {
	counter++
	discovered[here] = counter
	top++
	stack[top] = here
	ret := discovered[here]
	for i := 0; i < len(adj[here]); i++ {
		there := adj[here][i]
		if discovered[there] == 0 {
			ret = min(ret, tarjanSCC(there))
		} else if !inSCC[there] {
			ret = min(ret, discovered[there])
		}
	}

	if ret == discovered[here] {
		scc := newGroup()
		for {
			value := stack[top]
			top--
			scc.add(value)
			inSCC[value] = true
			if value == here {
				break
			}
		}
		sccGroups = append(sccGroups, scc)
	}

	return ret
}

func newGroup() (scc *group) {
	scc = new(group)
	scc.min = 1<<31
	scc.vertexList = make([]int, 0, 10)
	return
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
	doTarjanSCC()

	wr := bufio.NewWriter(os.Stdout)
	sort.Sort(sccGroups)
	wr.WriteString(strconv.Itoa(len(sccGroups)))
	wr.WriteByte('\n')

	for i := range sccGroups {
		sort.Ints(sccGroups[i].vertexList)
		for _, v := range sccGroups[i].vertexList {
			wr.WriteString(strconv.Itoa(v))
			wr.WriteByte(' ')
		}
		wr.WriteString("-1\n")
	}
	wr.Flush()
}
