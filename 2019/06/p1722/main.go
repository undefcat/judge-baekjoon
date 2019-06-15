package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc *bufio.Scanner
	fact [21]int
)

func init() {
	fact[0] = 1
	for i := 1; i <= 20; i++ {
		fact[i] = fact[i-1] * i
	}

	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

var (
	N    int
	perm []int
)

func main() {
	N = scanInt()
	subProblem := scanInt()

	if subProblem == 1 {
		solveOne()
	} else {
		solveTwo()
	}
}

var (
	skip int
	used []bool
)

// k번째의 순열 찾기
func solveOne() {
	skip = scanInt() - 1
	used = make([]bool, N+1)
	perm = make([]int, 0, N)

	generate(0, skip)
}

func generate(idx, skip int) {
	if idx == N {
		// print permutation
		for _, v := range perm {
			fmt.Printf("%d ", v)
		}
		return
	}

	for n := 1; n <= N; n++ {
		if used[n] {
			continue
		}

		if fact[N-idx-1] <= skip {
			skip -= fact[N-idx-1]
		} else {
			used[n] = true
			perm = append(perm, n)
			generate(idx+1, skip)
			return
		}
	}
}

// 해당 순열이 몇 번째인지 찾기
func solveTwo() {
	perm = make([]int, N)
	for pi := 0; pi < N; pi++ {
		perm[pi] = scanInt()
	}

	ans := 1
	for i := 0; i < N; i++ {
		less := 0
		for j := i + 1; j < N; j++ {
			if perm[j] < perm[i] {
				less++
			}
		}

		ans += fact[N-i-1] * less
	}

	fmt.Print(ans)
}
