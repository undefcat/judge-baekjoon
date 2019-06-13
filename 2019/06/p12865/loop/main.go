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
	N, K := scanInt(), scanInt()
	cache := make([]int, K+1)

	for i := 0; i < N; i++ {
		w, v := scanInt(), scanInt()
		for j := K; j >= w; j-- {
			if cache[j-w]+v > cache[j] {
				cache[j] = cache[j-w] + v
			}
		}
	}

	fmt.Print(cache[K])
}
