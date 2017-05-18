package main

import (
	"fmt"
	"github.com/tk-shirasaka/ansi"
	"github.com/tk-shirasaka/reversi/game"
)

func input() (int, int) {
	input := func(message string) int {
		num := 0

		for {
			fmt.Printf(ansi.Cursor(1, ansi.PREV_LINE) + ansi.Cursor(2, ansi.LINE_CLEAR))
			fmt.Printf(message)
			fmt.Scanf("%d", &num)
			if num >= 1 && num <= 8 {
				break
			}
		}

		return num
	}
	i := input("Input Line No : ") - 1
	j := input("Input Column No : ") - 1

	return i, j
}

func main() {
	field := game.Init()

	for {
		fmt.Println(field.String())
		field.Select(input())
		fmt.Printf(ansi.Clear())
	}
}
