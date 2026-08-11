package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_repository_instanceCmd = &cobra.Command{
	Use:   "instance",
	Short: "Instance management",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_repository_instanceCmd).Standalone()

	help_repositoryCmd.AddCommand(help_repository_instanceCmd)
}
