package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var editSourcesCmd = &cobra.Command{
	Use:   "edit-sources [file]",
	Short: "edit the source information file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(editSourcesCmd).Standalone()

	rootCmd.AddCommand(editSourcesCmd)

	carapace.Gen(editSourcesCmd).PositionalCompletion(
		carapace.ActionFiles().Chdir("/etc/apt/sources.list.d"),
	)
}
