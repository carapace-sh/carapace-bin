package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_repository_dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dump repository state information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_repository_dumpCmd).Standalone()

	help_repositoryCmd.AddCommand(help_repository_dumpCmd)
}
