package main

import "fmt"

func main() {
	var s string
	fmt.Scanf("%s", &s)

	minus := 0
	plus := 0
	isMinus := false
	num := 0
	beforeOperator := true

	for _, v := range s {
		switch v {
		case '-':
			if !isMinus {
				plus += num
				isMinus = true
			} else {
				isMinus = true
				minus += num
			}

			num = 0
			beforeOperator = true

		case '+':
			if isMinus {
				minus += num
			} else {
				plus += num
			}

			num = 0
			beforeOperator = true

		default:
			if beforeOperator && v == '0' {
				continue
			}

			beforeOperator = false
			num *= 10
			num += int(v-'0')
		}
	}
	if isMinus {
		minus += num
	} else {
		plus += num
	}
	fmt.Print(plus-minus)
}
