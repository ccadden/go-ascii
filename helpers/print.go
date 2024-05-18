package helpers

import "fmt"

type Char interface {
	byte | int
}

func PrintTable() {
	PrintHeader()
	for i := 33; i < 126; i++ {
		PrintRow(i)
	}
	fmt.Println()
	PrintSep()
	fmt.Println()
}

func PrintHeader() {
	PrintSep()
	fmt.Printf("\n|%5s|%5s|%5s|%9s|%5s|\n", "SYM", "DEC", "HEX", "BIN", "OCT")
	PrintSep()
}

func PrintRow[T Char](i T) {
	fmt.Printf("\n|%5c|%5v|%5X| %08b|%5o|", i, i, i, i, i)
}

func PrintSep() {
	fmt.Printf("|-----|-----|-----|---------|-----|")
}
