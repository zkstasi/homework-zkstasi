package main

import "fmt"

func main() {

	str1 := "#"
	str2 := " "
	num := 8

	var chessboard string

	for i := 0; i < num; i++ {

		emptyStr := ""

		for j := 0; j < num; j++ {

			if (i+j)%2 == 0 {
				emptyStr += str1
			} else {
				emptyStr += str2
			}
		}
		chessboard += emptyStr + "\n"
	}

	fmt.Println(chessboard)
}
