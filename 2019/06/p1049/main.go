package main

import (
	"bufio"
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

func main() {
	N, M := scanInt(), scanInt()

	packPrice, singlePrice := 1001, 1001

	for mi := 0; mi < M; mi++ {
		packPrice = min(packPrice, scanInt())
		singlePrice = min(singlePrice, scanInt())
	}

	// 만약 패키지의 가격이 줄보다 비싸면 걍 줄로 사는게 가장 쌈
	if packPrice >= singlePrice*6 {
		fmt.Print(N * singlePrice)
		return
	}

	// 최소한 패키지가 줄보다 싸므로
	// 일단 모두 패키지로 사본다.
	packAmount := (N-1)/6+1

	// 패키지로만 다 샀을때와 정확히 맞춰 샀을 때
	// 1. 살 것이 6의 배수면 패키지로만 사는게 정답
	allPackPrice := packAmount*packPrice
	if N%6 == 0 {
		fmt.Print(allPackPrice)
		return
	}
	// 2. 아닌 경우 패키지가 넘치므로
	fmt.Print(min(allPackPrice, (N%6)*singlePrice + allPackPrice - packPrice))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
