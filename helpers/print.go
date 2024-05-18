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
}

func PrintHeader() {
	fmt.Printf("|------|------|------|--------|--------|")
	fmt.Printf("\n|%6s|%6s|%6s|%8s|%8s|", "SYM", "DEC", "HEX", "BIN", "OCT")
	fmt.Printf("\n|------|------|------|--------|--------|")
}

func PrintRow[T Char](i T) {
	fmt.Printf("\n|%6c|%6v|%6X|%8b|%8o|", i, i, i, i, i)
}

func PrintFooter() {
	fmt.Printf("\n|------|------|------|--------|--------|\n")
}
