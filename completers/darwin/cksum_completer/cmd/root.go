package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cksum",
	Short: "display file checksums and count of bytes in a file",
	Long:  "https://keith.github.io/xcode-manpages/cksum.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("o", "o", "", "Use historical checksum format 1, 2, or 3")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
