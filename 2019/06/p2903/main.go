package main

import "fmt"

func main() {
	var N int
	fmt.Scanf("%d", &N)

	a := 2
	for i := 0; i < N; i++ {
		a = a*2-1
	}

	fmt.Print(a*a)
}