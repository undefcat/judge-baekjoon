package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Meeting struct {
	start, end int
}

type Meetings []*Meeting

func (it Meetings) Len() int {
	return len(it)
}

func (it Meetings) Less(i, j int) bool {
	if it[i].end == it[j].end {
		return it[i].start < it[j].start
	}

	return it[i].end < it[j].end
}

func (it Meetings) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
}

var (
	sc = bufio.NewScanner(os.Stdin)
	meetings Meetings
)

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	n := scanInt()
	meetings = make(Meetings, n)

	for i := 0; i < n; i++ {
		start, end := scanInt(), scanInt()
		meeting := &Meeting{start, end}
		meetings[i] = meeting
	}

	sort.Sort(meetings)

	before := &Meeting{-1, -1}
	count := 0
	for i := 0; i < n; i++ {
		if before.end == meetings[i].end {
			if meetings[i].start == meetings[i].end {
				count++
			}
			continue
		}

		if before.end <= meetings[i].start {
			before = meetings[i]
			count++
		}
	}
	fmt.Print(count)
}
