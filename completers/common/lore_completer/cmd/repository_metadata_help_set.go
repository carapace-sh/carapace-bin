package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_metadata_help_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set metadata on the repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_metadata_help_setCmd).Standalone()

	repository_metadata_helpCmd.AddCommand(repository_metadata_help_setCmd)
}
