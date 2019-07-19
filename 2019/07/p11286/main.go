package main

import (
	"bufio"
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

type Data struct {
	abs, num int
}

type Heap struct {
	arr  []Data
	size int
}

func NewHeap(size int) *Heap {
	return &Heap{
		arr: make([]Data, size),
		size: 0}
}

func (it *Heap) Push(x int) {
	it.arr[it.size] = Data{abs(x), x}
	it.size++
	it.up()
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func (it *Heap) up() {
	cur := it.size-1
	p := (cur-1)/2

	for cur > 0 {
		if it.arr[cur].abs > it.arr[p].abs {
			break
		}

		if it.arr[cur].abs == it.arr[p].abs &&
			it.arr[cur].num > it.arr[p].num {
			break
		}

		it.arr[cur], it.arr[p] = it.arr[p], it.arr[cur]

		cur = p
		p = (cur-1)/2
	}
}

func (it *Heap) Pop() int {
	if it.size == 0 {
		return 0
	}

	ret := it.arr[0]
	it.size--
	it.arr[0] = it.arr[it.size]
	it.down()

	return ret.num
}

func (it *Heap) down() {
	cur := 0

	for {
		l, r := cur*2+1, cur*2+2

		if l >= it.size {
			return
		}

		next := cur
		if it.arr[next].abs > it.arr[l].abs {
			next = l
		} else if it.arr[next].abs == it.arr[l].abs {
			if it.arr[next].num > it.arr[l].num {
				next = l
			}
		}

		if r < it.size {
			if it.arr[next].abs > it.arr[r].abs {
				next = r
			} else if it.arr[next].abs == it.arr[r].abs {
				if it.arr[next].num > it.arr[r].num {
					next = r
				}
			}
		}

		if cur == next {
			return
		}

		it.arr[cur], it.arr[next] = it.arr[next], it.arr[cur]
		cur = next
	}
}

func main() {
	n := scanInt()
	heap := NewHeap(n)

	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	for i := 0; i < n; i++ {
		cmd := scanInt()
		if cmd != 0 {
			heap.Push(cmd)
		} else {
			wr.WriteString(strconv.Itoa(heap.Pop()))
			wr.WriteByte('\n')
		}
	}
}
