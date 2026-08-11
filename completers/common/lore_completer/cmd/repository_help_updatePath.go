package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_help_updatePathCmd = &cobra.Command{
	Use:   "update-path",
	Short: "Update the stored path for this instance",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_help_updatePathCmd).Standalone()

	repository_helpCmd.AddCommand(repository_help_updatePathCmd)
}
