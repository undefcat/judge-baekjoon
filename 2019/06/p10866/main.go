package main

import (
	"bufio"
	"container/list"
	"os"
	"strconv"
)

var (
	sc *bufio.Scanner
	wr *bufio.Writer
	deque = list.New()
)

func init() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	wr = bufio.NewWriter(os.Stdout)
}

func scanText() string {
	sc.Scan()
	return sc.Text()
}

func scanInt() int {
	n, _ := strconv.Atoi(scanText())
	return n
}

func scanCommand() {
	switch scanText() {
	case "push_back":
		deque.PushBack(scanInt())
		return

	case "push_front":
		deque.PushFront(scanInt())
		return

	case "pop_front":
		e := deque.Front()
		if e == nil {
			wr.WriteString("-1")
		} else {
			deque.Remove(e)
			wr.WriteString(strconv.Itoa(e.Value.(int)))
		}

	case "pop_back":
		e := deque.Back()
		if e == nil {
			wr.WriteString("-1")
		} else {
			deque.Remove(e)
			wr.WriteString(strconv.Itoa(e.Value.(int)))
		}

	case "size":
		wr.WriteString(strconv.Itoa(deque.Len()))

	case "empty":
		if deque.Len() == 0 {
			wr.WriteByte('1')
		} else {
			wr.WriteByte('0')
		}

	case "front":
		e := deque.Front()
		if e == nil {
			wr.WriteString("-1")
		} else {
			wr.WriteString(strconv.Itoa(e.Value.(int)))
		}

	case "back":
		e := deque.Back()
		if e == nil {
			wr.WriteString("-1")
		} else {
			wr.WriteString(strconv.Itoa(e.Value.(int)))
		}
	}

	wr.WriteByte('\n')
}

func main() {
	defer wr.Flush()
	n := scanInt()
	for i := 0; i < n; i++ {
		scanCommand()
	}
}
