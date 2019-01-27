package main

import (
	"bufio"
	"os"
)

var (
	sc = bufio.NewScanner(os.Stdin)
)

func main() {
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()
	for {
		sc.Scan()
		input := sc.Bytes()
		if len(input) == 3 {
			if string(input) == "END" {
				return
			}
		}

		for i := len(input)-1; i >= 0; i-- {
			wr.WriteByte(input[i])
		}
		wr.WriteByte('\n')
	}
}
