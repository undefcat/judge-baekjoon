package bfs_binarySearch

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

type Adj struct {
	weight int
	city   int
}

var (
	N, M   int
	cities [][]*Adj
)

func main() {
	N, M = scanInt(), scanInt()

	cities = make([][]*Adj, N)
	for ci := range cities {
		cities[ci] = make([]*Adj, 0, 8)
	}

	for mi := 0; mi < M; mi++ {
		a, b, weight := scanInt()-1, scanInt()-1, scanInt()
		cities[a] = append(cities[a], &Adj{weight, b})
		cities[b] = append(cities[b], &Adj{weight, a})
	}

	from, to := scanInt()-1, scanInt()-1

	lo, hi := 0, 1000000000
	for lo <= hi {
		mid := (lo+hi)/2

		if bfs(from, to, mid) {
			lo = mid+1
		} else {
			hi = mid-1
		}
	}

	fmt.Println(hi)
}

func bfs(a, b, max int) bool {
	discovered := make([]bool, N)

	q := list.New()
	q.PushBack(a)

	for q.Len() != 0 {
		e := q.Front()
		q.Remove(e)

		here := e.Value.(int)
		if discovered[here] {
			continue
		}

		discovered[here] = true

		for _, city := range cities[here] {
			if discovered[city.city] || city.weight < max {
				continue
			}

			q.PushBack(city.city)
		}
	}

	return discovered[b]
}
