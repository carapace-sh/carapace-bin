package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var recompressCmd = &cobra.Command{
	Use:   "recompress",
	Short: "recompress an archive",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(recompressCmd).Standalone()

	rootCmd.AddCommand(recompressCmd)

	carapace.Gen(recompressCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
