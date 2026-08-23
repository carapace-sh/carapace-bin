package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xxd",
	Short: "make a hex dump or do the reverse",
	Long:  "https://keith.github.io/xcode-manpages/xxd.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("C", false, "Capitalize variable names in C include style")
	rootCmd.Flags().Bool("E", false, "Change character encoding to EBCDIC")
	rootCmd.Flags().String("R", "", "When to revert operation")
	rootCmd.Flags().Bool("autoskip", false, "Toggle autoskip: a single '*' replaces NUL-lines")
	rootCmd.Flags().BoolP("bits", "b", false, "Switch to bits (binary digits) dump")
	rootCmd.Flags().StringP("cols", "c", "", "Format octets per line")
	rootCmd.Flags().Bool("d", false, "Show offset in decimal instead of hex")
	rootCmd.Flags().Bool("e", false, "Switch to little-endian hex dump")
	rootCmd.Flags().StringP("groupsize", "g", "", "Separate output by bytes")
	rootCmd.Flags().BoolP("help", "h", false, "Print a summary of available commands")
	rootCmd.Flags().BoolP("include", "i", false, "Output in C include file style")
	rootCmd.Flags().StringP("len", "l", "", "Stop after writing len octets")
	rootCmd.Flags().StringP("name", "n", "", "Override the variable name output when -i is used")
	rootCmd.Flags().StringP("offset", "o", "", "Add offset to the displayed file position")
	rootCmd.Flags().BoolP("plain", "p", false, "Output in PostScript continuous hex dump style")
	rootCmd.Flags().BoolP("revert", "r", false, "Reverse operation: convert hex dump into binary")
	rootCmd.Flags().StringP("seek", "s", "", "Skip to the specified offset")
	rootCmd.Flags().Bool("u", false, "Use upper case hex letters")
	rootCmd.Flags().Bool("v", false, "Print version information")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"R": carapace.ActionValues("never", "always", "auto"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}