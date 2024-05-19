/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/ccadden/go-ascii/cmderrors"
	"github.com/spf13/cobra"
)

// charCmd represents the char command
var charCmd = &cobra.Command{
	Use:   "char",
	Short: "Returns a character for a given ASCII digit",
	Long: `Find corresponding ASCII character for a given digit
Supports translations for decimal (default), hex, binary, and octal.

Flags are mutually exclusive.`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			val, err := parseInput(cmd, arg)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("%c\n", byte(val))
			}

		}
	},
}

func parseInput(cmd *cobra.Command, arg string) (int64, error) {
	var val int64
	var err error

	decimal, _ := cmd.Flags().GetBool("decimal")
	hex, _ := cmd.Flags().GetBool("hex")
	bin, _ := cmd.Flags().GetBool("binary")
	oct, _ := cmd.Flags().GetBool("octal")

	if decimal || !(hex || bin || oct) {
		i, err := strconv.Atoi(arg)

		if err != nil {
			return 0, &cmderrors.CharError{
				Message:     "Problem parsing input as decimal",
				NumAsString: arg,
				Err:         err,
			}
		}

		return int64(i), nil

	} else if hex {
		val, err = strconv.ParseInt(arg, 16, 64)

		if err != nil {
			return 0, &cmderrors.CharError{
				Message:     "Problem parsing input as hex",
				NumAsString: arg,
				Err:         err,
			}
		}

	} else if bin {
		val, err = strconv.ParseInt(arg, 2, 64)

		if err != nil {
			return 0, &cmderrors.CharError{
				Message:     "Problem parsing input as binary",
				NumAsString: arg,
				Err:         err,
			}
		}

	} else if oct {
		val, err = strconv.ParseInt(arg, 8, 64)

		if err != nil {
			return 0, &cmderrors.CharError{
				Message:     "Problem parsing input as octal",
				NumAsString: arg,
				Err:         err,
			}
		}

	}

	return val, nil
}

func init() {
	rootCmd.AddCommand(charCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// charCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// charCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	charCmd.Flags().BoolP("decimal", "d", false, "Reads input as decimal value (default).")
	charCmd.Flags().BoolP("hex", "x", false, "Reads input as hexidecimal value.")
	charCmd.Flags().BoolP("binary", "b", false, "Reads input as a binary value.")
	charCmd.Flags().BoolP("octal", "o", false, "Reads input as an octal value.")

	charCmd.MarkFlagsMutuallyExclusive("decimal", "hex", "binary", "octal")
}
