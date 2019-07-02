package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	if n%2 != 0 {
		fmt.Print(0)
		return
	}

	cache := make([]int, n+1)
	cache[0] = 1
	cache[2] = 3

	for i := 4; i <= n; i += 2 {
		// 기본적으로 2칸을 차지하는 블럭의 종류는 3가지이다.
		cache[i] = cache[i-2]*3

		// 2칸이 증가할 때마다 새로운 모양의 블럭을 2가지 놓을 수 있다.
		// 따라서 예를 들어 10인 경우
		// 10짜리 2개
		// 8짜리 2개놓기 * 2짜리 놓는 경우의 수
		// 6짜리 2개놓기 * 4짜리 놓는 경우의 수
		// ...
		// 이런식으로 계산하면 된다.
		for j := 4; i-j >= 0; j += 2 {
			cache[i] += cache[i-j]*2
		}
	}

	fmt.Print(cache[n])
}