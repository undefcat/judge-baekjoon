package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	sc     *bufio.Scanner
	X, Y   int
	matrix [][]bool
	target [][]bool
)

func scanYX() (int, int) {
	sc.Scan()
	input := strings.Fields(sc.Text())

	y, _ := strconv.Atoi(input[0])
	x, _ := strconv.Atoi(input[1])

	return y, x
}

func scanX() []bool {
	sc.Scan()
	input := sc.Bytes()

	x := make([]bool, X)
	for i := range input {
		if input[i] == '1' {
			x[i] = true
		}
	}

	return x
}

func toggle(y, x int) {
	for dy := 0; dy < 3; dy++ {
		for dx := 0; dx < 3; dx++ {
			matrix[y+dy][x+dx] = !matrix[y+dy][x+dx]
		}
	}
}

func check() bool {
	for y := 0; y < Y; y++ {
		for x := 0; x < X; x++ {
			if matrix[y][x] != target[y][x] {
				return false
			}
		}
	}
	return true
}

func init() {
	sc = bufio.NewScanner(os.Stdin)
}

func main() {
	Y, X = scanYX()
	matrix = make([][]bool, Y)
	for i := 0; i < Y; i++ {
		matrix[i] = scanX()
	}

	target = make([][]bool, Y)
	for i := 0; i < Y; i++ {
		target[i] = scanX()
	}

	ans := 0
	for y := 0; y <= Y-3; y++ {
		for x := 0; x <= X-3; x++ {
			if matrix[y][x] != target[y][x] {
				toggle(y, x)
				ans++
			}
		}
	}

	if check() {
		fmt.Print(ans)
	} else {
		fmt.Print(-1)
	}
}
