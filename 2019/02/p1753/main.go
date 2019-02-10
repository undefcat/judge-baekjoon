package main

import (
	"bufio"
	"container/heap"
	"os"
	"strconv"
)

type Edge struct {
	vertex int
	weight int
}

type Edges map[int]*Edge

var (
	sc      *bufio.Scanner
	V, E, K int
	adj     []Edges
)

func nextInt() int {
	sc.Scan()
	b := sc.Bytes()

	ret := 0
	for _, v := range b {
		ret *= 10
		ret += int(v) - 48
	}

	return ret
}

type Distance struct {
	vertex int
	cost   int
}

type PriorityQueue []*Distance

func (it PriorityQueue) Len() int {
	return len(it)
}

func (it PriorityQueue) Less(i, j int) bool {
	return it[i].cost < it[j].cost
}

func (it PriorityQueue) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
}

func (it *PriorityQueue) Push(x interface{}) {
	*it = append(*it, x.(*Distance))
}

// heap.Pop을 수행하면, 맨 위의 요소가 맨 마지막으로 옮겨지고
// heap.Interface의 Pop을 호출한다.
// 따라서 마지막 요소를 리턴하고, 큐의 크기를 줄이는 구현을 하면 된다.
func (it *PriorityQueue) Pop() interface{} {
	l := it.Len() - 1
	old := *it
	item := old[l]
	*it = old[:l]
	return item
}

const INF = 1 << 60

func dijkstra() (dist []int) {
	dist = make([]int, V+1)
	dist[0] = INF

	for i := 1; i < len(dist); i *= 2 {
		copy(dist[i:], dist[:i])
	}

	dist[K] = 0
	pq := make(PriorityQueue, 0, V+1)
	heap.Push(&pq, &Distance{K, 0})

	for pq.Len() != 0 {
		distance := heap.Pop(&pq).(*Distance)
		from, c := distance.vertex, distance.cost

		if dist[from] < c {
			continue
		}

		for to, edge := range adj[from] {
			cost := c + edge.weight
			if dist[to] > cost {
				dist[to] = cost
				heap.Push(&pq, &Distance{to, cost})
			}
		}
	}

	return
}

func init() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
}

func main() {
	V, E, K = nextInt(), nextInt(), nextInt()

	adj = make([]Edges, V+1)
	for i := range adj {
		adj[i] = make(Edges)
	}

	for i := 0; i < E; i++ {
		u, v, w := nextInt(), nextInt(), nextInt()

		edge, ok := adj[u][v]
		if ok && edge.weight < w {
			continue
		}

		adj[u][v] = &Edge{v, w}
	}

	dist := dijkstra()

	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	for i := 1; i <= V; i++ {
		if dist[i] < INF {
			wr.WriteString(strconv.Itoa(dist[i]))
		} else {
			wr.WriteString("INF")
		}
		wr.WriteByte('\n')
	}
}
