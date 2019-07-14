package main

import "fmt"

var (
	N                    int
	X                    []bool
	board                [][]bool
	rDiagonalMap         [][]int
	lDiagonalMap         [][]int
	rDiagonal, lDiagonal []bool
)

func dfs(y int) int {
	if y == N {
		return 1
	}

	ret := 0
	for x := 0; x < N; x++ {
		if X[x] {
			continue
		}

		ld, rd := lDiagonalMap[y][x], rDiagonalMap[y][x]

		if lDiagonal[ld] || rDiagonal[rd] {
			continue
		}

		board[y][x] = true
		lDiagonal[ld] = true
		rDiagonal[rd] = true
		X[x] = true

		ret += dfs(y + 1)

		board[y][x] = false
		lDiagonal[ld] = false
		rDiagonal[rd] = false
		X[x] = false
	}

	return ret
}

func main() {
	fmt.Scanf("%d", &N)

	X = make([]bool, N)
	board = make([][]bool, N)


	rDiagonalMap = make([][]int, N)
	for ri := range rDiagonalMap {
		rDiagonalMap[ri] = make([]int, N)
	}
	// rDiagonal
	for x := N - 1; x >= 0; x-- {
		for y := 0; y < N; y++ {
			rDiagonalMap[y][x] = N - 1 - x + y
		}
	}

	lDiagonalMap = make([][]int, N)
	for li := range lDiagonalMap {
		lDiagonalMap[li] = make([]int, N)
	}

	// yDiagonal
	for x := 0; x < N; x++ {
		for y := 0; y < N; y++ {
			lDiagonalMap[y][x] = x+y
		}
	}

	rDiagonal = make([]bool, 2*N-1)
	lDiagonal = make([]bool, 2*N-1)

	for bi := range board {
		board[bi] = make([]bool, N)
	}

	fmt.Println(dfs(0))
}
