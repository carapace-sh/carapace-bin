package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "Index the specified files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(indexCmd).Standalone()

	rootCmd.AddCommand(indexCmd)

	carapace.Gen(indexCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
	carapace.Gen(indexCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
