package main

import "fmt"

const MOD = 9901

func main() {
	var n int
	fmt.Scanf("%d", &n)

	a, b, c := 1, 1, 1
	for i := 1; i < n; i++ {
		d := (a+b+c)%MOD
		e := (a+c)%MOD
		f := (a+b)%MOD
		a, b, c = d, e, f
	}

	fmt.Print((a+b+c)%MOD)
}
