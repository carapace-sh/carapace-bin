package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "shasum",
	Short: "print or check SHA checksums",
	Long:  "https://keith.github.io/xcode-manpages/shasum.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("0", "0", false, "Read NUL-terminated filenames")
	rootCmd.Flags().BoolS("U", "U", false, "Universal mode")
	rootCmd.Flags().BoolS("b", "b", false, "Binary mode")
	rootCmd.Flags().BoolS("c", "c", false, "Check checksums")
	rootCmd.Flags().BoolS("h", "h", false, "Display help")
	rootCmd.Flags().BoolS("q", "q", false, "Quiet mode")
	rootCmd.Flags().BoolS("s", "s", false, "Skip checksum verification")
	rootCmd.Flags().BoolS("t", "t", false, "Text mode (default)")
	rootCmd.Flags().BoolS("v", "v", false, "Verbose")
	rootCmd.Flags().BoolS("w", "w", false, "Warn about checksum format")

	rootCmd.Flags().StringS("a", "a", "", "Algorithm (1, 224, 256, 384, 512)")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"a": carapace.ActionValues("1", "224", "256", "384", "512", "512224", "512256"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
