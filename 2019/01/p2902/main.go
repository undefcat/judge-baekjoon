package main

import (
	"bufio"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	text := sc.Bytes()

	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	wr.WriteByte(text[0])
	text = text[1:]

	for i := range text {
		if text[i] == '-' {
			wr.WriteByte(text[i+1])
		}
	}
}