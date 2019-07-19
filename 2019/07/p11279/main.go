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

type MaxHeap struct {
	arr  []int
	size int
}

func NewMaxHeap(size int) *MaxHeap {
	return &MaxHeap{
		arr: make([]int, size),
		size: 0}
}

func (it *MaxHeap) Push(x int) {
	it.arr[it.size] = x
	it.size++
	it.up()
}

func (it *MaxHeap) up() {
	cur := it.size-1
	p := (cur-1)/2

	for cur > 0 && it.arr[cur] > it.arr[p] {
		it.arr[cur], it.arr[p] = it.arr[p], it.arr[cur]
		cur = p
		p = (cur-1)/2
	}
}

func (it *MaxHeap) Pop() int {
	if it.size == 0 {
		return 0
	}

	ret := it.arr[0]
	it.size--
	it.arr[0] = it.arr[it.size]
	it.down()
	return ret
}

func (it *MaxHeap) down() {
	cur := 0

	for {
		l, r := cur*2+1, cur*2+2

		if l >= it.size {
			return
		}

		next := cur
		if it.arr[next] < it.arr[l] {
			next = l
		}

		if r < it.size && it.arr[next] < it.arr[r] {
			next = r
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
	heap := NewMaxHeap(n)

	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	for i := 0; i < n; i++ {
		cmd := scanInt()
		if cmd > 0 {
			heap.Push(cmd)
		} else {
			wr.WriteString(strconv.Itoa(heap.Pop()))
			wr.WriteByte('\n')
		}
	}
}
