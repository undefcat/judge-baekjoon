package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	a, b, tmp := big.NewInt(0), big.NewInt(1), big.NewInt(0)

	for i := 0; i < n; i++ {
		tmp.Add(a, b)
		a.Set(b)
		b.Set(tmp)
	}

	fmt.Print(a)
}
