package main

import (
	"bufio"
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
	ret := []int{0, 0, 0}

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

type tree struct {
	size int
	t []int
}

func New(arr []int) *tree {
	t := &tree{
		size: len(arr),
		t: make([]int, len(arr)*4),
	}

	t.init(arr, 0, t.size-1, 1)
	return t
}

func (it *tree) init(arr []int, l, r, node int) int {
	ret := &it.t[node]
	if l == r {
		*ret = arr[l]
		return *ret
	}

	mid := (l+r)/2
	left := it.init(arr, l, mid, node*2)
	right := it.init(arr, mid+1, r, node*2+1)
	*ret = left+right

	return *ret
}

func (it *tree) query(l, r, node, ln, lr int) int {
	if r < ln || l > lr {
		return 0
	}

	if l <= ln && lr <= r {
		return it.t[node]
	}

	mid := (ln+lr)/2

	return it.query(l, r, node*2, ln, mid) +
		it.query(l, r, node*2+1, mid+1, lr)
}

func (it *tree) Query(l, r int) int {
	return it.query(l, r, 1, 0, it.size-1)
}

func (it *tree) update(idx, value, node, nl, nr int) int {
	ret := &it.t[node]
	if idx < nl || idx > nr {
		return *ret
	}

	if nl == nr {
		*ret = value
		return *ret
	}

	mid := (nl+nr)/2
	*ret = it.update(idx, value, node*2, nl, mid) +
		it.update(idx, value, node*2+1, mid+1, nr)

	return *ret
}

func (it *tree) Update(idx, value int) int {
	return it.update(idx, value, 1, 0, it.size-1)
}

var (
	rd *bufio.Reader
	wr *bufio.Writer
)

func init() {
	rd = bufio.NewReaderSize(os.Stdin, 1<<15)
	wr = bufio.NewWriterSize(os.Stdout, 1<<10)
}

func main() {
	defer wr.Flush()

	input := scanInts()
	N, M, K := input[0], input[1], input[2]

	nums := make([]int, N)
	for i := 0; i < N; i++ {
		nums[i] = scanInt()
	}

	t := New(nums)

	MK := M+K
	for i := 0; i < MK; i++ {
		input := scanInts()
		if input[0] == 1 {
			t.Update(input[1]-1, input[2])
		} else {
			result := t.Query(input[1]-1, input[2]-1)
			wr.WriteString(strconv.Itoa(result))
			wr.WriteByte('\n')
		}
	}
}