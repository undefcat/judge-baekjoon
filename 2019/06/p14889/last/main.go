package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func init() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
}

func scan() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

var (
	n     int
	ans   = math.MaxInt32
	power [][]int
	t1    []int
	t2    []int
	t1p   int
	t2p   int
)

func solve(x int) {
	if ans == 0 {
		return
	}

	if x == n {
		calc()
		return
	}

	// team1을 먼저 다 뽑고
	if t1p < len(t1) {
		t1[t1p] = x
		t1p++
		solve(x+1)
		t1p--
	}

	// team2를 다 뽑는다
	if t2p < len(t2) {
		t2[t2p] = x
		t2p++
		solve(x+1)
		t2p--
	}
}

func calc() {
	sum := 0

	for i := 0; i < len(t1)-1; i++ {
		for j := i+1; j < len(t1); j++ {
			sum += power[t1[i]][t1[j]] + power[t1[j]][t1[i]]
			sum -= power[t2[i]][t2[j]] + power[t2[j]][t2[i]]
		}
	}

	if sum < 0 {
		sum *= -1
	}

	if ans > sum {
		ans = sum
	}

	return
}

func main() {
	n = scan()

	power = make([][]int, n)
	for i := range power {
		power[i] = make([]int, n)
		for j := range power[i] {
			power[i][j] = scan()
		}
	}

	t1 = make([]int, n/2)
	t2 = make([]int, n/2)

	solve(0)
	fmt.Print(ans)
}
