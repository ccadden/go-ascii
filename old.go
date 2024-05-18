package main

import (
	"fmt"
	"os"
)

type Char interface {
	byte | int
}

func main() {
	if os.Args[1] == "table" {
		printTable()
	} else {
		printHeader()
		for _, arg := range os.Args[1:] {
			for _, char := range []byte(arg) {
				printRow(char)
			}
		}

		printFooter()
	}
}

func printTable() {
	printHeader()
	for i := 33; i < 126; i++ {
		printRow(i)
	}
	fmt.Println()
}

func printHeader() {
	fmt.Printf("|------|------|------|--------|--------|")
	fmt.Printf("\n|%6s|%6s|%6s|%8s|%8s|", "SYM", "DEC", "HEX", "BIN", "OCT")
	fmt.Printf("\n|------|------|------|--------|--------|")
}

func printRow[T Char](i T) {
	fmt.Printf("\n|%6c|%6v|%6X|%8b|%8o|", i, i, i, i, i)
}

func printFooter() {
	fmt.Printf("\n|------|------|------|--------|--------|\n")
}
