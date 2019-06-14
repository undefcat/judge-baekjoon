package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc    *bufio.Scanner
	antic int
)

var alpha = []byte{
	'b', 'd', 'e', 'f', 'g',
	'h', 'j', 'k', 'l', 'm',
	'o', 'p', 'q', 'r', 's',
	'u', 'v', 'w', 'x', 'y', 'z',
}

func init() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	for _, v := range "antic" {
		antic |= 1 << uint(v-'a')
	}

	for i := range alpha {
		alpha[i] -= 'a'
	}
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
	N, K  int
	words []int
	ans   int
)

func main() {
	N, K = scanInt(), scanInt()

	if K < 5 {
		fmt.Print(0)
		return
	}

	if K == 26 {
		fmt.Print(N)
		return
	}

	words = make([]int, N)
	for ni := 0; ni < N; ni++ {
		words[ni] = stringToBit(scanString())
	}

	pick(antic, 0, K-5)
	fmt.Print(ans)
}

func stringToBit(str string) int {
	ret := 0
	for _, v := range str {
		ret |= 1 << uint(v-'a')
	}
	return ret
}

func existChar(word, set int) bool {
	return (word & set) == word
}

func pick(bit, cur, remain int) {
	if remain == 0 {
		sum := 0
		for i := range words {
			if existChar(words[i], bit) {
				sum++
			}
		}

		if sum > ans {
			ans = sum
		}

		return
	}

	for i := cur; i < 21; i++ {
		nextBit := bit | 1<<uint(alpha[i])
		pick(nextBit, i+1, remain-1)
	}
}
