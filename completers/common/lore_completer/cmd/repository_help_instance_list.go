package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_help_instance_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all registered instances for this repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_help_instance_listCmd).Standalone()

	repository_help_instanceCmd.AddCommand(repository_help_instance_listCmd)
}
