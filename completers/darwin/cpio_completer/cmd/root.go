package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cpio",
	Short: "copy files to and from archives",
	Long:  "https://keith.github.io/xcode-manpages/cpio.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("0", "0", false, "Read filenames separated by NUL")
	rootCmd.Flags().BoolS("6", "6", false, "PWB format")
	rootCmd.Flags().BoolS("A", "A", false, "Append to existing archive")
	rootCmd.Flags().BoolS("B", "B", false, "Block at 5120 bytes per record")
	rootCmd.Flags().BoolS("H", "H", false, "Format")
	rootCmd.Flags().BoolS("I", "I", false, "Input file")
	rootCmd.Flags().BoolS("J", "J", false, "xz compression")
	rootCmd.Flags().BoolS("L", "L", false, "Follow symbolic links")
	rootCmd.Flags().BoolS("O", "O", false, "Output file")
	rootCmd.Flags().BoolS("V", "V", false, "Verbose")
	rootCmd.Flags().BoolS("a", "a", false, "Reset access times")
	rootCmd.Flags().BoolS("b", "b", false, "Swap bytes and halfwords")
	rootCmd.Flags().BoolS("d", "d", false, "Create directories as needed")
	rootCmd.Flags().BoolS("f", "f", false, "Pattern matching")
	rootCmd.Flags().BoolS("i", "i", false, "Input mode")
	rootCmd.Flags().BoolS("l", "l", false, "Link files")
	rootCmd.Flags().BoolS("m", "m", false, "Preserve modification time")
	rootCmd.Flags().BoolS("n", "n", false, "Numeric UID and GID")
	rootCmd.Flags().BoolS("o", "o", false, "Output mode")
	rootCmd.Flags().BoolS("p", "p", false, "Pass-through mode")
	rootCmd.Flags().BoolS("q", "q", false, "Quiet")
	rootCmd.Flags().BoolS("r", "r", false, "Rename files")
	rootCmd.Flags().BoolS("t", "t", false, "List contents")
	rootCmd.Flags().BoolS("u", "u", false, "Unconditional")
	rootCmd.Flags().BoolS("v", "v", false, "Verbose")
	rootCmd.Flags().BoolS("z", "z", false, "gzip compression")
}
