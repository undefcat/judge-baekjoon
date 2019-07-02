package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer
)

func init() {
	rd = bufio.NewReaderSize(os.Stdin, 1<<20)
	wr = bufio.NewWriterSize(os.Stdout, 1<<15)
}

func scanInt() int {
	l, _, _ := rd.ReadLine()
	ret := 0
	for _, c := range l {
		ret *= 10
		ret += int(c-'0')
	}
	return ret
}

func scanInts() (int, int) {
	l, _, _ := rd.ReadLine()
	ret := [2]int{}

	i := 0
	for _, c := range l {
		if c == ' ' {
			i++
			continue
		}

		ret[i] *= 10
		ret[i] += int(c-'0')
	}

	return ret[0], ret[1]
}

func main() {
	defer wr.Flush()

	n, m := scanInts()
	arr := make([]int, n)
	for ni := 0; ni < n; ni++ {
		arr[ni] = scanInt()
	}

	t := newTree(arr)
	for mi := 0; mi < m; mi++ {
		l, r := scanInts()
		wr.WriteString(strconv.Itoa(t.Query(l-1, r-1)))
		wr.WriteByte('\n')
	}
}

type tree struct {
	last int
	t []int
}

func newTree(arr []int) *tree {
	ret := &tree{
		last: len(arr)-1,
		t: make([]int, len(arr)*4),
	}

	ret.init(arr, 0, ret.last, 1)
	return ret
}

func (it *tree) init(arr []int, l, r, n int) int {
	ret := &it.t[n]
	if l == r {
		*ret = arr[l]
		return *ret
	}

	mid := (l+r)/2

	*ret = min(it.init(arr, l, mid, n*2),
		it.init(arr, mid+1, r, n*2+1))

	return *ret
}

const Max = (1<<63)-1

func (it *tree) query(l, r, n, nl, nr int) int {
	if r < nl || l > nr {
		return Max
	}

	if l <= nl && nr <= r {
		return it.t[n]
	}

	mid := (nl+nr)/2

	return min(it.query(l, r, n*2, nl, mid),
		it.query(l, r, n*2+1, mid+1, nr))
}

func (it *tree) Query(l, r int) int {
	return it.query(l, r, 1, 0, it.last)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}