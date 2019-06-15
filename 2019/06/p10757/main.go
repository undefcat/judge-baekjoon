package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

var (
	rd *bufio.Reader
	wr *bufio.Writer
)

func init() {
	rd = bufio.NewReaderSize(os.Stdin, 21111)
	wr = bufio.NewWriterSize(os.Stdout, 21111)
}

func readBytes() []byte {
	buf, _ := rd.ReadBytes(' ')
	return buf
}

func main() {
	var bigA, bigB big.Int

	bigA.SetString(string(readBytes()), 10)
	bigB.SetString(string(readBytes()), 10)

	fmt.Print(bigA.Add(&bigA, &bigB))
}