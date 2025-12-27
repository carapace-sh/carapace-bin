package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dlopenMetadataCmd = &cobra.Command{
	Use:   "dlopen-metadata",
	Short: "Parse and print ELF dlopen metadata",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dlopenMetadataCmd).Standalone()

	rootCmd.AddCommand(dlopenMetadataCmd)

	carapace.Gen(dlopenMetadataCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
