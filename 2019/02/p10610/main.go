package main

import (
	"bufio"
	"os"
	"sort"
)

type b []byte

func (it b) Len() int {
	return len(it)
}

func (it b) Less(i, j int) bool {
	return it[i] > it[j]
}

func (it b) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
}

func main() {
	rd := bufio.NewReaderSize(os.Stdin, 111111)
	input, _, _ := rd.ReadLine()

	existZero := false
	sum := 0
	for _, v := range input {
		if v == '0' {
			existZero = true
		}
		sum += int(v-'0')
	}

	wr := bufio.NewWriter(os.Stdout)
	// 10의 배수가 되려면 0이 존재해야하고
	// 3의 배수가 되려면 각 자리 숫자의 합이 3의 배수여야 한다.
	if existZero && sum%3 == 0 {
		// 내림차순 정렬하고
		// 출력하면 된다.
		sort.Sort(b(input))
		wr.Write(input)
	} else {
		wr.WriteString("-1")
	}
	wr.Flush()
}