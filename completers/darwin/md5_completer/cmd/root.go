package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "md5",
	Short: "calculate a message-digest fingerprint (checksum) for a file",
	Long:  "https://keith.github.io/xcode-manpages/md5.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("p", "p", false, "Echo stdin to stdout and compute checksum")
	rootCmd.Flags().BoolS("q", "q", false, "Quiet mode")
	rootCmd.Flags().BoolS("r", "r", false, "Reverse format")
	rootCmd.Flags().BoolS("t", "t", false, "Test mode")
	rootCmd.Flags().BoolS("x", "x", false, "Extended format")

	rootCmd.Flags().StringS("c", "c", "", "Compare checksum to string")
	rootCmd.Flags().StringS("s", "s", "", "Compute checksum of string")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
