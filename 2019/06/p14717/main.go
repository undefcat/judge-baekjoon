package main

import "fmt"

var (
	a1, a2 int
	cards  []int
)

func isWin(b1, b2 int) bool {
	a, b := isDdang(a1, a2), isDdang(b1, b2)

	if a && b {
		return a1 > b1
	}

	if a && !b {
		return true
	}

	if !a && b {
		return false
	}

	return (a1+a2)%10 > (b1+b2)%10
}

func isDdang(a, b int) bool {
	return a == b
}

func main() {
	fmt.Scanf("%d %d", &a1, &a2)

	cards = make([]int, 11)
	for i := 1; i <= 10; i++ {
		cards[i] = 2
	}

	cards[a1]--
	cards[a2]--

	win := 0

	for i := 1; i <= 10; i++ {
		b1 := 0
		if cards[i] == 0 {
			continue
		}
		b1 = i
		cards[i]--
		for j := i; j <= 10; j++ {
			b2 := 0
			if cards[j] == 0 {
				continue
			}
			b2 = j
			cards[j]--

			if isWin(b1, b2) {
				if cards[i] == 0 && cards[j] == 0 {
					win += 1
				} else if cards[i] == 0 || cards[j] == 0 {
					win += 2
				} else {
					win += 4
				}
			}

			cards[j]++
		}
		cards[i]++
	}

	ans := float64(win)/float64(9*17)
	fmt.Printf("%.3f", ans)
}
