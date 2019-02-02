package main

import "fmt"

func main() {
	var N, M, K int
	fmt.Scanf("%d %d %d", &N, &M, &K)

	team := 0
	n := N/2

	if n > M {
		team = M
		N -= 2*team
		M = 0
	} else {
		team = n
		N -= 2*n
		M -= team
	}

	remain := N+M
	if K > remain {
		need := K-remain
		team -= (need-1)/3 + 1
	}

	fmt.Print(team)
}
