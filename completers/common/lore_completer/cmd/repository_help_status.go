package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_help_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show current repository status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_help_statusCmd).Standalone()

	repository_helpCmd.AddCommand(repository_help_statusCmd)
}
