package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_repositoryCmd = &cobra.Command{
	Use:   "repository",
	Short: "Repository commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_repositoryCmd).Standalone()

	helpCmd.AddCommand(help_repositoryCmd)
}
