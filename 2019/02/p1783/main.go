package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var x, y int
	fmt.Scanf("%d %d", &y, &x)

	x--

	switch y {
	case 1:
		fmt.Print(1)

	case 2:
		ans := x/2+1
		if ans >= 5 {
			fmt.Print(4)
		} else {
			fmt.Print(ans)
		}

	default:
		ans := x+1
		if x < 6 {
			fmt.Print(min(4, ans))
		} else {
			fmt.Print(ans-2)
		}
	}
}
