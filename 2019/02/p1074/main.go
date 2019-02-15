package main

import (
	"fmt"
	"os"
)

var (
	N     int
	r, c  int
	ans   int
)

func visit(y, x, size int) {
	if (r < y || r > y+size) && (c < x || c > x+size) {
		ans += size*size
		return
	}

	if size == 2 {
		for dy := 0; dy < 2; dy++ {
			for dx := 0; dx < 2; dx++ {
				if y+dy == r && x+dx == c {
					fmt.Println(ans)
					os.Exit(0)
				}
				ans++
			}
		}
		return
	}

	nextSize := size >> 1
	yn, xn := y+nextSize, x+nextSize

	visit(y, x, nextSize)
	visit(y, xn, nextSize)
	visit(yn, x, nextSize)
	visit(yn, xn, nextSize)
}

func main() {
	fmt.Scanf("%d %d %d", &N, &r, &c)
	visit(0, 0, 1<<uint(N))
}
