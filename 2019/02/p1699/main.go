package main

import (
	"fmt"
	"math"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	dp := make([]int, n+1)
	dp[0] = 1<<63-1

	for i := 1; i < len(dp); i *= 2 {
		copy(dp[i:], dp[:i])
	}

	end := int(math.Sqrt(float64(n)))

	for i := 1; i <= end; i++ {
		dp[i*i] = 1
	}

	for i := 2; i <= n; i++ {
		if dp[i] == 1 {
			continue
		}

		sqrtI := int(math.Sqrt(float64(i)))
		for j := sqrtI; j >= 1; j-- {
			square := j*j
			dp[i] = min(dp[i], dp[square] + dp[i-square])
		}
	}

	fmt.Print(dp[n])
}
