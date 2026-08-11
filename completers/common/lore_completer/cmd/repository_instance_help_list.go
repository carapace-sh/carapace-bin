package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_instance_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all registered instances for this repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_instance_help_listCmd).Standalone()

	repository_instance_helpCmd.AddCommand(repository_instance_help_listCmd)
}
