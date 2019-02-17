package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	sc		  *bufio.Scanner
	wr        *bufio.Writer
)

func f(in, post []int) {
	if len(post) == 0 {
		return
	}

	root := post[0]
	wr.WriteString(strconv.Itoa(root))
	wr.WriteByte(' ')

	post = post[1:]
	leftIn, rightIn := split(in, root)
	rightInLen := len(rightIn)

	rightPost := post[:rightInLen]
	leftPost := post[rightInLen:]
	f(leftIn, leftPost)
	f(rightIn, rightPost)
}

func split(in []int, rootValue int) ([]int, []int) {
	inrootIdx := -1
	for i := range in {
		if in[i] == rootValue {
			inrootIdx = i
			break
		}
	}

	leftTree, rightTree := in[:inrootIdx], in[inrootIdx+1:]
	return leftTree, rightTree
}

func scanInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	wr = bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	n := scanInt()
	inorder := make([]int, n)
	for i := 0; i < n; i++ {
		inorder[i] = scanInt()
	}

	postorder := make([]int, n)
	for i := n-1; i >= 0; i-- {
		postorder[i] = scanInt()
	}

	f(inorder, postorder)
}
