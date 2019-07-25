package main

import (
	"bufio"
	"bytes"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

type Marble struct {
	y, x int
}

const (
	Wall = byte('#')
	Out  = byte('O')
)

type State struct {
	beforeDir int
	count     int
	red, blue Marble
}

func (it *State) IsOverlap() bool {
	return (it.red.y == it.blue.y) && (it.red.x == it.blue.x)
}

func (it *State) Opposite(dir int) bool {
	switch it.beforeDir {
	case Up:
		return dir == Down

	case Down:
		return dir == Up

	case Left:
		return dir == Right

	case Right:
		return dir == Left

	}

	return false
}

var (
	board [][]byte
	N, M  int
)

const (
	Up = iota
	Right
	Down
	Left
)

var delta = [4][2]int{
	{-1, 0}, // UP
	{0, 1},  // Right
	{1, 0},  // Down
	{0, -1}, // Left
}

// 더 움직일 수 있으면 true, 없으면 false
func (it *Marble) Move(dir int) bool {
	it.y += delta[dir][0]
	it.x += delta[dir][1]

	if it.y >= N || it.y < 0 || it.x >= M || it.x < 0 || board[it.y][it.x] == Wall {
		it.y -= delta[dir][0]
		it.x -= delta[dir][1]
		return false
	}

	if board[it.y][it.x] == Out {
		return false
	}

	return true
}

func (it *Marble) IsOut() bool {
	return board[it.y][it.x] == Out
}

// 현재 State를 dir방향으로 기울인 뒤
// red, blue의 위치를 리턴한다.
// 더 이상 게임이 진행이 안되는 경우 false를 리턴한다.
// 더 이상 게임이 진행이 안되는 경우는
// 1. red가 탈출했다.
// 2. blue가 탈출했다.
// 3. count가 10번째인데 red가 탈출하지 못했다.
// 4. 둘 다 움직이지 않았다.
func (it *State) NextState(dir int) (State, bool) {
	next := *it
	next.count++
	next.beforeDir = dir

	// 움직이지 않을 때까지 기울인다.
	moveCount := 0
	for next.red.Move(dir) {
		moveCount++
	}

	for next.blue.Move(dir) {
		moveCount++
	}

	// 만약 둘 다 움직이지 않았다면
	if moveCount == 0 {
		return next, false
	}

	// 둘 중 하나라도 탈출했다면
	if next.red.IsOut() || next.blue.IsOut() {
		return next, false
	}

	// 10번째인 경우
	if next.count == 10 {
		return next, false
	}

	// 겹쳐있는지 확인한다.
	if !next.IsOverlap() {
		return next, true
	}

	// 겹쳐있으면 보정한다.
	switch dir {
	case Up:
		if it.red.y < it.blue.y {
			next.blue.y++
		} else {
			next.red.y++
		}

	case Down:
		if it.red.y > it.blue.y {
			next.blue.y--
		} else {
			next.red.y--
		}

	case Left:
		if it.red.x < it.blue.x {
			next.blue.x++
		} else {
			next.red.x++
		}

	case Right:
		if it.red.x > it.blue.x {
			next.blue.x--
		} else {
			next.red.x--
		}
	}

	return next, true
}

var sc *bufio.Scanner

func init() {
	sc = bufio.NewScanner(os.Stdin)
}

func scanNM() {
	sc.Scan()
	split := bytes.Fields(sc.Bytes())

	N, _ = strconv.Atoi(string(split[0]))
	M, _ = strconv.Atoi(string(split[1]))
}

func scanLine() []byte {
	sc.Scan()
	buf := sc.Bytes()
	ret := make([]byte, len(buf))
	copy(ret, buf)
	return ret
}

func main() {
	scanNM()
	board = make([][]byte, N)

	for ni := 0; ni < N; ni++ {
		board[ni] = scanLine()
	}

	q := list.New()
	q.PushBack(StartState())

	for q.Len() != 0 {
		e := q.Front()
		q.Remove(e)
		here := e.Value.(State)

		for dir := 0; dir < 4; dir++ {
			if here.Opposite(dir) {
				continue
			}

			next, ok := here.NextState(dir)

			if !ok {
				if next.red.IsOut() && !next.blue.IsOut() {
					fmt.Println(next.count)
					return
				}

				continue
			}

			q.PushBack(next)
		}
	}

	fmt.Println("-1")
}

func StartState() State {
	ret := State{beforeDir: -1}

	for y := 0; y < N; y++ {
		for x := 0; x < M; x++ {
			if board[y][x] == 'R' {
				ret.red = Marble{y, x}
				board[y][x] = '.'
			}

			if board[y][x] == 'B' {
				ret.blue = Marble{y, x}
				board[y][x] = '.'
			}
		}
	}

	return ret
}