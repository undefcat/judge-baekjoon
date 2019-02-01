package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	sc  = bufio.NewScanner(os.Stdin)
	minus []int
	plus []int
)

func scanInt() int {
	sc.Scan()
	sum := 0
	minus := 1
	input := sc.Bytes()
	if input[0] == '-' {
		minus = -1
		input = input[1:]
	}

	for _, v := range input {
		sum *= 10
		sum += int(v) - '0'
	}

	return sum * minus
}

func main() {
	n := scanInt()
	minus = make([]int, 0, n)
	plus = make([]int, 0, n)
	isZero := false
	ans := 0

	for i := 0; i < n; i++ {
		n := scanInt()
		switch {
		case n > 1:
			plus = append(plus, n)

		case n < 0:
			minus = append(minus, n)

		default:
			ans += n
			if n == 0 {
				isZero = true
			}
		}
	}

	sort.Ints(plus)
	sort.Ints(minus)

	if len(plus)%2 == 1 {
		ans += plus[0]
		plus = plus[1:]
	}

	if len(minus)%2 == 1 {
		if !isZero {
			ans += minus[len(minus)-1]
		}

		minus = minus[:len(minus)-1]
	}

	for i := 1; i < len(plus); i += 2 {
		ans += plus[i-1]*plus[i]
	}

	for i := 1; i < len(minus); i += 2 {
		ans += minus[i-1]*minus[i]
	}

	fmt.Print(ans)
}
