/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/ccadden/go-ascii/helpers"
	"github.com/spf13/cobra"
)

// lookupCmd represents the lookup command
var lookupCmd = &cobra.Command{
	Use:   "lookup",
	Short: "Look up ASCII values",
	Long: `A small program for querying ASCII values in a variety of formats

Decimal:
ascii lookup abcd
49, 50, 51, 52

Hex:
ascii lookup -x abcd
31, 32, 33, 34

Binary and Octal:
ascii lookup -o -b abcd
01100001, 01100010, 01100011, 01100100
141, 142, 143, 144

For multiple flags, print priority is:
1) Decimal
2) Hex
3) Binary
4) Octal
`,
	Run: func(cmd *cobra.Command, args []string) {
		table, _ := cmd.Flags().GetBool("table")

		for _, arg := range args {
			var d []string
			var h []string
			var b []string
			var o []string

			if table {
				helpers.PrintHeader()
			}
			for _, char := range []byte(arg) {
				if table {
					helpers.PrintRow(char)
				} else {
					hex, _ := cmd.Flags().GetBool("hex")
					bin, _ := cmd.Flags().GetBool("binary")
					oct, _ := cmd.Flags().GetBool("octal")

					if hex {
						h = append(h, fmt.Sprintf("%X", char))
					}

					if bin {
						b = append(b, fmt.Sprintf("%08b", char))
					}

					if oct {
						o = append(o, fmt.Sprintf("%o", char))
					}

					if decimal, _ := cmd.Flags().GetBool("decimal"); decimal || !(hex || bin || oct) {
						d = append(d, fmt.Sprintf("%v", char))
					}

				}
			}

			if table {
				fmt.Println()
				helpers.PrintSep()
				fmt.Println()
			} else {
				if len(d) > 0 {
					fmt.Println(strings.Join(d, ", "))
				}

				if len(h) > 0 {
					fmt.Println(strings.Join(h, ", "))
				}

				if len(b) > 0 {
					fmt.Println(strings.Join(b, ", "))
				}

				if len(o) > 0 {
					fmt.Println(strings.Join(o, ", "))
				}
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(lookupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lookupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lookupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	lookupCmd.Flags().BoolP("decimal", "d", false, "Print decimal values for inputs (default).")
	lookupCmd.Flags().BoolP("hex", "x", false, "Print hexidecimal values for inputs.")
	lookupCmd.Flags().BoolP("binary", "b", false, "Print binary values for inputs.")
	lookupCmd.Flags().BoolP("octal", "o", false, "Print octal values for inputs.")
	lookupCmd.Flags().BoolP("table", "t", false, "Print table for given inputs. Overwrites other flags.")
}
