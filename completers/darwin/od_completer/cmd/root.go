package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "od",
	Short: "octal, decimal, hex, ASCII dump",
	Long:  "https://keith.github.io/xcode-manpages/od.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("B", "B", false, "Octal bytes")
	rootCmd.Flags().BoolS("D", "D", false, "Decimal longs")
	rootCmd.Flags().BoolS("F", "F", false, "Floating-point longs")
	rootCmd.Flags().BoolS("H", "H", false, "Hexadecimal longs")
	rootCmd.Flags().BoolS("I", "I", false, "Hexadecimal longs")
	rootCmd.Flags().BoolS("L", "L", false, "Decimal longs")
	rootCmd.Flags().BoolS("O", "O", false, "Octal longs")
	rootCmd.Flags().BoolS("X", "X", false, "Hexadecimal longs")
	rootCmd.Flags().BoolS("a", "a", false, "Named characters")
	rootCmd.Flags().BoolS("b", "b", false, "Octal bytes")
	rootCmd.Flags().BoolS("c", "c", false, "Characters")
	rootCmd.Flags().BoolS("d", "d", false, "Decimal shorts")
	rootCmd.Flags().BoolS("e", "e", false, "Floating-point longs")
	rootCmd.Flags().BoolS("f", "f", false, "Floating-point shorts")
	rootCmd.Flags().BoolS("h", "h", false, "Hexadecimal shorts")
	rootCmd.Flags().BoolS("i", "i", false, "Hexadecimal shorts")
	rootCmd.Flags().BoolS("l", "l", false, "Decimal longs")
	rootCmd.Flags().BoolS("o", "o", false, "Octal shorts")
	rootCmd.Flags().BoolS("s", "s", false, "Decimal shorts")
	rootCmd.Flags().BoolS("v", "v", false, "Do not suppress duplicates")
	rootCmd.Flags().BoolS("x", "x", false, "Hexadecimal shorts")

	rootCmd.Flags().StringS("A", "A", "", "Address format")
	rootCmd.Flags().StringS("N", "N", "", "Number of bytes")
	rootCmd.Flags().StringS("j", "j", "", "Skip bytes")
	rootCmd.Flags().StringS("t", "t", "", "Type")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
