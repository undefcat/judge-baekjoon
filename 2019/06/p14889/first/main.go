package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func init() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
}

func scan() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

var picked []bool
func solve(current, remain int) {
	if remain == 0 {
		calc()
		return
	}

	if current == len(picked) {
		return
	}

	for cur := current; cur < len(picked); cur++ {
		picked[cur] = true
		solve(cur+1, remain-1)
		picked[cur] = false
		// 두 팀이므로 어떤 한 요소가 포함된 모든 팀을 다 구하면
		// 나머지는 결정되어서 더 돌 필요가 없다.
		// 이 경우 맨 첫 선수를 포함한 모든 팀을 고려하면
		// 첫 선수가 포함되지 않은 다음 조합들은
		// 결국 이전에 구한 팀의 반대팀이 된다.
        if cur == 0 {
            return
        }
	}
}

var s [][]int
var ans = math.MaxInt32

func calc() {
	sum := 0
	for i := 0; i < len(picked); i++ {
		flag := 1
		if picked[i] {
			flag = -1
		}

		for j := 0; j < len(picked); j++ {
			if i == j {
				continue
			}

			if picked[i] == picked[j] {
				sum += s[i][j]*flag
			}
		}
	}

	ans = min(ans, abs(sum))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func main() {
	n := scan()
	s = make([][]int, n)
	picked = make([]bool, n)
	for i := range s {
		s[i] = make([]int, n)
		for j := range s[i] {
			s[i][j] = scan()
		}
	}

	solve(0, n/2)
	fmt.Print(ans)
}