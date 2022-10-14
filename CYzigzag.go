package main

import "fmt"

func main() {

	zigzag(4, "HELLOWORLDPEOPLE")

}

func zigzag(baris int, huruf string) {
	if baris > 0 {
		longprint := len(huruf) - 1
		fmt.Println(longprint)
		baris = baris - 1

		pa := 0
		for a := 0; a <= baris; a++ {
			Za := 0

			for i := a; i <= longprint; i = i + 2 {

				//fmt.Print(pa)
				if (Za+pa)%baris == 0 {
					fmt.Print(string(huruf[i]), "  ")
					Za++

				} else if (Za % baris) == 0 {
					fmt.Print(string(huruf[i]), "  ")
					Za++

				} else {
					fmt.Print("   ")
					Za++
				}

			}
			fmt.Println("")
			pa++

		}

	}
}
