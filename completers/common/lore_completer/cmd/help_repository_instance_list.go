package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_repository_instance_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all registered instances for this repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_repository_instance_listCmd).Standalone()

	help_repository_instanceCmd.AddCommand(help_repository_instance_listCmd)
}
