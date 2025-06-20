package main

import "fmt"

func homeWork3(count int) {

	str1 := "#"
	str2 := " "

	var chessboard string

	for i := 0; i < count; i++ {

		emptyStr := ""

		for j := 0; j < count; j++ {

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
