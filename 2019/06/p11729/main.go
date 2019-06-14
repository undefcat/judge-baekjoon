package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	wr = bufio.NewWriterSize(os.Stdout, (1<<22)+100)
	count int
)

func move(n, left, mid, right byte) {
	if n == 1 {
		wr.WriteByte(left)
		wr.WriteByte(' ')
		wr.WriteByte(right)
		wr.WriteByte('\n')
		count++
		return
	}

	if n == 2 {
		wr.WriteByte(left)
		wr.WriteByte(' ')
		wr.WriteByte(mid)
		wr.WriteByte('\n')
		wr.WriteByte(left)
		wr.WriteByte(' ')
		wr.WriteByte(right)
		wr.WriteByte('\n')
		wr.WriteByte(mid)
		wr.WriteByte(' ')
		wr.WriteByte(right)
		wr.WriteByte('\n')
		count += 3
		return
	}

	move(n-1, left, right, mid)
	move(1, left, mid, right)
	move(n-1, mid, left, right)
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	move(byte(n), '1', '2', '3')
	fmt.Println(count)
	wr.Flush()
}
