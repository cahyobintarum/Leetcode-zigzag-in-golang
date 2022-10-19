package main

import "fmt"

func main() {

	convert("HELLOWORLDPEOPLE", 4)

}

func convert(s string, numRows int) {
	if numRows > 0 {
		longprint := len(s) - 1
		fmt.Println(longprint)
		numRows = numRows - 1

		move := 0
		for a := 0; a <= numRows; a++ {
			position := 0

			for i := a; i <= longprint; i = i + 2 {

				//fmt.Print(move)
				if (position+move)%numRows == 0 {
					fmt.Print(string(s[i]), "  ")
					position++

				} else if (position % numRows) == 0 {
					fmt.Print(string(s[i]), "  ")
					position++

				} else {
					fmt.Print("   ")
					position++
				}

			}
			fmt.Println(" ")
			move++

		}

	}
}
