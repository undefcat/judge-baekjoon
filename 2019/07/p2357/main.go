package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

func scanInt() int {
	l, _, _ := rd.ReadLine()
	ret := 0
	for _, c := range l {
		ret *= 10
		ret += int(c-'0')
	}
	return ret
}

func scanInts() []int {
	l, _, _ := rd.ReadLine()
	ret := []int{0, 0}

	i := 0
	for _, c := range l {
		if c == ' ' {
			i++
			continue
		}

		ret[i] *= 10
		ret[i] += int(c-'0')
	}

	return ret
}

type minmax struct {
	min, max int
}

type tree struct {
	size int
	t []minmax
}

func New(arr []int) *tree {
	ret := &tree{
		size: len(arr),
		t: make([]minmax, len(arr)*4),
	}

	ret.init(arr, 0, ret.size-1, 1)
	return ret
}

func (it *tree) init(arr []int, l, r, node int) minmax {
	ret := &it.t[node]
	if l == r {
		*ret = minmax{
			min: arr[l],
			max: arr[l],
		}

		return *ret
	}

	mid := (l+r)/2

	*ret = merge(it.init(arr, l, mid, node*2),
		it.init(arr, mid+1, r, node*2+1))

	return *ret
}

func merge(a, b minmax) minmax {
	return minmax{
		min: min(a.min, b.min),
		max: max(a.max, b.max),
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (it *tree) query(l, r, node, nl, nr int) minmax {
	if r < nl || l > nr {
		return minmax{
			min: math.MaxInt64,
			max: -1,
		}
	}

	if l <= nl && nr <= r {
		return it.t[node]
	}

	mid := (nl+nr)/2
	return merge(it.query(l, r, node*2, nl, mid),
		it.query(l, r, node*2+1, mid+1, nr))
}

func (it *tree) Query(l, r int) minmax {
	return it.query(l, r, 1, 0, it.size-1)
}

var (
	rd *bufio.Reader
	wr *bufio.Writer
)

func init() {
	rd = bufio.NewReaderSize(os.Stdin, 1<<10)
	wr = bufio.NewWriterSize(os.Stdout, 1<<10)
}

func main() {
	input := scanInts()
	arr := make([]int, input[0])

	for i := range arr {
		arr[i] = scanInt()
	}

	t := New(arr)
	for i := 0; i < input[1]; i++ {
		lr := scanInts()
		result := t.Query(lr[0]-1, lr[1]-1)

		wr.WriteString(strconv.Itoa(result.min))
		wr.WriteByte(' ')
		wr.WriteString(strconv.Itoa(result.max))
		wr.WriteByte('\n')
	}

	wr.Flush()
}