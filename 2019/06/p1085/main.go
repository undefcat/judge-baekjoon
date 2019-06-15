package main

import "fmt"

func main() {
	var x, y, w, h int

	fmt.Scanf("%d %d %d %d", &x, &y, &w, &h)

	ans := min(x, abs(w-x))
	ans = min(ans, abs(h-y))
	ans = min(ans, y)

	fmt.Print(ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}