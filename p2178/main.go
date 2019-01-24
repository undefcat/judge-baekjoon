package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

var (
	sc    = bufio.NewScanner(os.Stdin)
	N, M  int
	maze  [][]byte
	visit [][]bool
)

type node struct {
	x, y     int
	distance int
}

func init() {
	sc.Split(bufio.ScanWords)
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func main() {
	N, M = scanInt(), scanInt()
	maze = make([][]byte, N)
	visit = make([][]bool, N)
	for i := range maze {
		maze[i] = make([]byte, M)
		visit[i] = make([]bool, M)
		sc.Scan()
		input := sc.Bytes()
		for j := range maze[i] {
			maze[i][j] = input[j]
		}
	}

	root := &node{0, 0, 1}
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() != 0 {
		e := queue.Front()
		queue.Remove(e)
		n := e.Value.(*node)

		if n.y == N-1 && n.x == M-1 {
			fmt.Print(n.distance)
			return
		}

		// 위로 간다.
		if n.y > 0 && maze[n.y-1][n.x] == '1' && !visit[n.y-1][n.x] {
			queue.PushBack(&node{n.x, n.y-1, n.distance+1})
			visit[n.y-1][n.x] = true
		}

		// 아래로
		if n.y+1 < N && maze[n.y+1][n.x] == '1' && !visit[n.y+1][n.x] {
			queue.PushBack(&node{n.x, n.y+1, n.distance+1})
			visit[n.y+1][n.x] = true
		}

		// 왼쪽
		if n.x > 0 && maze[n.y][n.x-1] == '1' && !visit[n.y][n.x-1] {
			queue.PushBack(&node{n.x-1, n.y, n.distance+1})
			visit[n.y][n.x-1] = true
		}

		//오른쪽
		if n.x+1< M && maze[n.y][n.x+1] == '1' && !visit[n.y][n.x+1] {
			queue.PushBack(&node{n.x+1, n.y, n.distance+1})
			visit[n.y][n.x+1] = true
		}
	}
}
