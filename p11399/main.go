package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
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
	times := make([]int, n)
	for i := 0 ; i < n; i++ {
		times[i] = scanInt()
	}

	sort.Ints(times)

	ans := 0
	before := 0
	for i := range times {
		ans += before + times[i]
		before += times[i]
	}
	fmt.Print(ans)
}
