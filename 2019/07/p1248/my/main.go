package my

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func init() {
	sc = bufio.NewScanner(os.Stdin)
}

func scanString() string {
	sc.Scan()
	return sc.Text()
}

func scanInt() int {
	n, _ := strconv.Atoi(scanString())
	return n
}

var (
	N    int
	sign [][]byte
	pick []int
)

func main() {
	N = scanInt()

	sign = make([][]byte, N)
	for si := range sign {
		sign[si] = make([]byte, N)
	}

	str := scanString()
	for i := 0; i < N; i++ {
		for j := i; j < N; j++ {
			sign[i][j] = str[0]
			str = str[1:]
		}
	}

	pick = make([]int, N)

	switch sign[0][0] {
	case '+':
		for n := 1; n <= 10; n++ {
			pick[0] = n
			if ok := solve(1); ok {
				break
			}
		}

	case '0':
		pick[0] = 0
		solve(1)

	case '-':
		for n := -1; n >= -10; n-- {
			pick[0] = n
			if ok := solve(1); ok {
				break
			}
		}
	}

	for _, v := range pick {
		fmt.Printf("%d ", v)
	}
}

func solve(i int) bool {
	if i == N {
		return true
	}

	switch sign[i][i] {
	case '+':
		start := 1
		if pick[i-1] < 0 && sign[i-1][i] == '+' {
			start = abs(pick[i-1])+1
		}

		for n := start; n <= 10; n++ {
			pick[i] = n
			if !check(i) {
				continue
			}

			if ok := solve(i+1); ok {
				return true
			}
		}

	case '0':
		pick[i] = 0

		if ok := solve(i+1); ok {
			return true
		}

	case '-':
		start := -1
		if pick[i-1] > 0 && sign[i-1][i] == '-' {
			start = abs(pick[i-1])+1
		}

		for n := start; n >= -10; n-- {
			pick[i] = n
			if !check(i) {
				continue
			}

			if ok := solve(i+1); ok {
				return true
			}
		}
	}

	return false
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func check(i int) bool {
	ans := 0
	for n := i; n >= 0; n-- {
		ans += pick[n]
		switch sign[n][i] {
		case '+':
			if ans <= 0 {
				return false
			}

		case '0':
			if ans != 0 {
				return false
			}

		case '-':
			if ans >= 0 {
				return false
			}
		}
	}

	return true
}
