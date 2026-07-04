package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "typeset",
	Short: "Define variables and set attributes",
	Long:  "https://zsh.sourceforge.io/Doc/Release/Shell-Builtin-Commands.html#index-typeset",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("A", "A", false, "the names refer to associative array parameters")
	rootCmd.Flags().BoolS("E", "E", false, "use floating point with scientific notation")
	rootCmd.Flags().BoolS("F", "F", false, "use floating point with fixed point notation")
	rootCmd.Flags().BoolS("H", "H", false, "hide value: typeset will not display the value of the parameter")
	rootCmd.Flags().BoolS("L", "L", false, "left justify and remove leading blanks")
	rootCmd.Flags().BoolS("R", "R", false, "right justify and fill with leading blanks")
	rootCmd.Flags().BoolS("T", "T", false, "tie a scalar and an array parameter together")
	rootCmd.Flags().BoolS("U", "U", false, "keep only the first occurrence of each value in the array")
	rootCmd.Flags().BoolS("Z", "Z", false, "zero fill and right justify")
	rootCmd.Flags().BoolS("a", "a", false, "the names refer to array parameters")
	rootCmd.Flags().BoolS("f", "f", false, "the names refer to functions rather than parameters")
	rootCmd.Flags().BoolS("g", "g", false, "the resulting parameter will not be made local")
	rootCmd.Flags().BoolS("h", "h", false, "hide: only useful for special parameters")
	rootCmd.Flags().BoolS("i", "i", false, "use integer arithmetic")
	rootCmd.Flags().BoolS("l", "l", false, "convert the result to lower case")
	rootCmd.Flags().BoolS("m", "m", false, "the name arguments are taken as patterns")
	rootCmd.Flags().BoolS("p", "p", false, "print the attributes and values of each name")
	rootCmd.Flags().BoolS("r", "r", false, "make readonly")
	rootCmd.Flags().BoolS("t", "t", false, "mark as tagged")
	rootCmd.Flags().BoolS("u", "u", false, "convert the result to upper case")
	rootCmd.Flags().BoolS("x", "x", false, "mark for automatic export to the environment")
}
