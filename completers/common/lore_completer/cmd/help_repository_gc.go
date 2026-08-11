package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_repository_gcCmd = &cobra.Command{
	Use:   "gc",
	Short: "Run a full garbage collection pass on the local repository store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_repository_gcCmd).Standalone()

	help_repositoryCmd.AddCommand(help_repository_gcCmd)
}
