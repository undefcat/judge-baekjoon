package main

import (
	"fmt"
)

func isValid(n int) bool {
	six := 0
	before := false
	for n != 0 {
		if n%10 == 6 {
			if before {
				six++
				if six == 3 {
					return true
				}
			} else {
				six = 1
			}
			before = true
		} else {
			before = false
			six = 0
		}
		n /= 10
	}
	return false
}

func main() {
	var target int
	fmt.Scanf("%d", &target)
	count := 0
	num := 666
	for {
		if isValid(num) {
			count++
			if count == target {
				fmt.Print(num)
				return
			}
		}
		num++
	}
}