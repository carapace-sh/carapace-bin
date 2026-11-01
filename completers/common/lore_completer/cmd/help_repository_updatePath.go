package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_repository_updatePathCmd = &cobra.Command{
	Use:   "update-path",
	Short: "Update the stored path for this instance",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_repository_updatePathCmd).Standalone()

	help_repositoryCmd.AddCommand(help_repository_updatePathCmd)
}
